package service

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"

	"backend/internal/models"
	"backend/internal/repository"
)

const (
	fontURL  = "https://github.com/google/fonts/raw/main/apache/roboto/Roboto-Regular.ttf" // URL для загрузки шрифта
	fontPath = "fonts/arial.ttf"                                                           // Локальный путь для сохранения шрифта
)

type DataService struct {
	dataRepo *repository.DataRepository
}

func NewDataService(repo *repository.DataRepository) *DataService {
	return &DataService{dataRepo: repo}
}

func (s *DataService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.dataRepo.GetUser(id.String())
}

// internal/service/data.go

func (s *DataService) GetGameStatsByCode(code string) ([]models.GameUser, error) {
	return s.dataRepo.GetGameUsersByCode(code)
}

// GenerateCertificate формирует PDF сертификат для пользователя
func (s *DataService) GenerateCertificate(userID uuid.UUID) ([]byte, error) {
	// Проверяем и скачиваем шрифт, если он отсутствует
	err := downloadFontIfNotExists(fontURL, fontPath)
	if err != nil {
		log.Printf("❌ Ошибка при загрузке шрифта: %v", err)
		return nil, fmt.Errorf("ошибка загрузки шрифта: %v", err)
	}

	// Получаем данные пользователя
	user, err := s.dataRepo.GetGameUserByID(userID)
	if err != nil {
		log.Printf("❌ Ошибка при получении пользователя: %v", err)
		return nil, err
	}

	// Получаем данные об игре
	game, err := s.dataRepo.GetGameLinkByID(user.GameID)
	if err != nil {
		log.Printf("❌ Ошибка при получении данных игры: %v", err)
		return nil, err
	}

	// Инициализация PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("Arial", "", fontPath)
	pdf.SetFont("Arial", "", 14)
	pdf.AddPage()

	// Рисуем рамку
	drawBorder(pdf)

	// Заголовок
	pdf.SetFont("Arial", "", 24)
	pdf.SetY(40)
	pdf.CellFormat(190, 10, "СЕРТИФИКАТ", "0", 0, "C", false, 0, "")

	pdf.Ln(20)

	// Основное содержимое
	pdf.SetFont("Arial", "", 14)
	pdf.MultiCell(0, 10, fmt.Sprintf(`
Настоящий сертификат подтверждает, что

%s

прошёл интерактивный урок по пожарной безопасности "%s"

в школе №%s, класс %s.

Дата выдачи: %s
`,
		user.FullName,
		game.GameName,
		game.SchoolNum,
		game.Class,
		time.Now().Format("02.01.2006"),
	), "", "C", false)

	pdf.Ln(10)

	// Линия для подписи
	pdf.SetY(230)
	pdf.Line(40, 230, 160, 230)
	pdf.SetXY(40, 235)
	pdf.Cell(120, 10, "Подпись директора")

	// Генерация PDF
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		log.Printf("❌ Ошибка генерации PDF: %v", err)
		return nil, err
	}

	return buf.Bytes(), nil
}

// drawBorder рисует рамку вокруг сертификата
func drawBorder(pdf *gofpdf.Fpdf) {
	pdf.SetDrawColor(0, 0, 0) // Черная рамка
	pdf.SetLineWidth(1)
	pdf.Rect(5, 5, 200, 287, "D") // Внешняя рамка
	pdf.SetLineWidth(0.5)
	pdf.Rect(10, 10, 190, 277, "D") // Внутренняя рамка
}

// downloadFontIfNotExists проверяет, существует ли файл шрифта, и загружает его, если отсутствует
func downloadFontIfNotExists(url, filepath string) error {
	// Проверяем, существует ли файл
	if _, err := os.Stat(filepath); err == nil {
		return nil // Файл уже существует
	}

	log.Println("📥 Загрузка шрифта...")
	// Создаем директорию, если её нет
	dir := filepath[:len(filepath)-len("arial.ttf")]
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Загружаем файл шрифта
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создаем файл для сохранения шрифта
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Копируем содержимое в файл
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	log.Println("✅ Шрифт успешно загружен и сохранён!")
	return nil
}
