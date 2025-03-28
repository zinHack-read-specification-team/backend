package service

import (
	"bytes"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"

	"backend/internal/models"
	"backend/internal/repository"
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

func (s *DataService) GenerateCertificate(userID uuid.UUID) ([]byte, error) {
	user, err := s.dataRepo.GetGameUserByID(userID)
	if err != nil {
		return nil, err
	}

	game, err := s.dataRepo.GetGameLinkByID(user.GameID)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("Arial", "", "/app/fonts/arial.ttf") // Путь к шрифту
	pdf.SetFont("Arial", "", 14)
	pdf.AddPage()

	// Рисуем рамку для сертификата
	drawBorder(pdf)

	// Заголовок "СЕРТИФИКАТ"
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

Дата: %s
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

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// drawBorder рисует красивую рамку
func drawBorder(pdf *gofpdf.Fpdf) {
	pdf.SetDrawColor(0, 0, 0)
	pdf.SetLineWidth(1)
	pdf.Rect(5, 5, 200, 287, "D")
	pdf.SetLineWidth(0.5)
	pdf.Rect(10, 10, 190, 277, "D")
}
