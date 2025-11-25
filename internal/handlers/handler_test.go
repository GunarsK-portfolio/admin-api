package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/GunarsK-portfolio/admin-api/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// =============================================================================
// Test Constants
// =============================================================================

const (
	testCertName      = "AWS Solutions Architect"
	testCertIssuer    = "Amazon Web Services"
	testCertIssueDate = "2024-01-15"
	testCertBadgeID   = "ABC123XYZ"                               // #nosec G101 -- test data, not a real credential
	testCertBadgeURL  = "https://aws.amazon.com/verify/ABC123XYZ" // #nosec G101 -- test data, not a real credential
	testSkillName     = "Go"
	testSkillTypeName = "Programming Languages"
	testCompanyName   = "Acme Corp"
	testPosition      = "Senior Developer"
	testStartDate     = "2020-01-01"
)

// =============================================================================
// Mock Repository
// =============================================================================

type mockRepository struct {
	// Profile
	getProfileFunc          func(ctx context.Context) (*models.Profile, error)
	updateProfileFunc       func(ctx context.Context, profile *models.Profile) error
	updateProfileAvatarFunc func(ctx context.Context, fileID int64) error
	deleteProfileAvatarFunc func(ctx context.Context) error
	updateProfileResumeFunc func(ctx context.Context, fileID int64) error
	deleteProfileResumeFunc func(ctx context.Context) error

	// Work Experience
	getAllWorkExperienceFunc  func(ctx context.Context) ([]models.WorkExperience, error)
	getWorkExperienceByIDFunc func(ctx context.Context, id int64) (*models.WorkExperience, error)
	createWorkExperienceFunc  func(ctx context.Context, exp *models.WorkExperience) error
	updateWorkExperienceFunc  func(ctx context.Context, exp *models.WorkExperience) error
	deleteWorkExperienceFunc  func(ctx context.Context, id int64) error

	// Certifications
	getAllCertificationsFunc func(ctx context.Context) ([]models.Certification, error)
	getCertificationByIDFunc func(ctx context.Context, id int64) (*models.Certification, error)
	createCertificationFunc  func(ctx context.Context, cert *models.Certification) error
	updateCertificationFunc  func(ctx context.Context, cert *models.Certification) error
	deleteCertificationFunc  func(ctx context.Context, id int64) error

	// Miniature Themes
	getAllMiniatureThemesFunc func(ctx context.Context) ([]models.MiniatureTheme, error)
	getMiniatureThemeByIDFunc func(ctx context.Context, id int64) (*models.MiniatureTheme, error)
	createMiniatureThemeFunc  func(ctx context.Context, theme *models.MiniatureTheme) error
	updateMiniatureThemeFunc  func(ctx context.Context, theme *models.MiniatureTheme) error
	deleteMiniatureThemeFunc  func(ctx context.Context, id int64) error

	// Miniature Projects
	getAllMiniatureProjectsFunc func(ctx context.Context) ([]models.MiniatureProject, error)
	getMiniatureProjectByIDFunc func(ctx context.Context, id int64) (*models.MiniatureProject, error)
	createMiniatureProjectFunc  func(ctx context.Context, project *models.MiniatureProject) error
	updateMiniatureProjectFunc  func(ctx context.Context, project *models.MiniatureProject) error
	deleteMiniatureProjectFunc  func(ctx context.Context, id int64) error
	addImageToProjectFunc       func(ctx context.Context, miniatureFile *models.MiniatureFile) error

	// Miniature Paints
	getAllMiniaturePaintsFunc func(ctx context.Context) ([]models.MiniaturePaint, error)
	getMiniaturePaintByIDFunc func(ctx context.Context, id int64) (*models.MiniaturePaint, error)
	createMiniaturePaintFunc  func(ctx context.Context, paint *models.MiniaturePaint) error
	updateMiniaturePaintFunc  func(ctx context.Context, paint *models.MiniaturePaint) error
	deleteMiniaturePaintFunc  func(ctx context.Context, id int64) error

	// Skills
	getAllSkillsFunc func(ctx context.Context) ([]models.Skill, error)
	getSkillByIDFunc func(ctx context.Context, id int64) (*models.Skill, error)
	createSkillFunc  func(ctx context.Context, skill *models.Skill) error
	updateSkillFunc  func(ctx context.Context, skill *models.Skill) error
	deleteSkillFunc  func(ctx context.Context, id int64) error

	// Skill Types
	getAllSkillTypesFunc func(ctx context.Context) ([]models.SkillType, error)
	getSkillTypeByIDFunc func(ctx context.Context, id int64) (*models.SkillType, error)
	createSkillTypeFunc  func(ctx context.Context, skillType *models.SkillType) error
	updateSkillTypeFunc  func(ctx context.Context, skillType *models.SkillType) error
	deleteSkillTypeFunc  func(ctx context.Context, id int64) error

	// Portfolio Projects
	getAllPortfolioProjectsFunc func(ctx context.Context) ([]models.PortfolioProject, error)
	getPortfolioProjectByIDFunc func(ctx context.Context, id int64) (*models.PortfolioProject, error)
	createPortfolioProjectFunc  func(ctx context.Context, project *models.PortfolioProject) error
	updatePortfolioProjectFunc  func(ctx context.Context, project *models.PortfolioProject) error
	deletePortfolioProjectFunc  func(ctx context.Context, id int64) error

	// Images/Files
	deleteImageFunc func(ctx context.Context, id int64) error
}

// Profile implementations
func (m *mockRepository) GetProfile(ctx context.Context) (*models.Profile, error) {
	if m.getProfileFunc != nil {
		return m.getProfileFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	if m.updateProfileFunc != nil {
		return m.updateProfileFunc(ctx, profile)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateProfileAvatar(ctx context.Context, fileID int64) error {
	if m.updateProfileAvatarFunc != nil {
		return m.updateProfileAvatarFunc(ctx, fileID)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteProfileAvatar(ctx context.Context) error {
	if m.deleteProfileAvatarFunc != nil {
		return m.deleteProfileAvatarFunc(ctx)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateProfileResume(ctx context.Context, fileID int64) error {
	if m.updateProfileResumeFunc != nil {
		return m.updateProfileResumeFunc(ctx, fileID)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteProfileResume(ctx context.Context) error {
	if m.deleteProfileResumeFunc != nil {
		return m.deleteProfileResumeFunc(ctx)
	}
	return errors.New("not implemented")
}

// Work Experience implementations
func (m *mockRepository) GetAllWorkExperience(ctx context.Context) ([]models.WorkExperience, error) {
	if m.getAllWorkExperienceFunc != nil {
		return m.getAllWorkExperienceFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetWorkExperienceByID(ctx context.Context, id int64) (*models.WorkExperience, error) {
	if m.getWorkExperienceByIDFunc != nil {
		return m.getWorkExperienceByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) CreateWorkExperience(ctx context.Context, exp *models.WorkExperience) error {
	if m.createWorkExperienceFunc != nil {
		return m.createWorkExperienceFunc(ctx, exp)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateWorkExperience(ctx context.Context, exp *models.WorkExperience) error {
	if m.updateWorkExperienceFunc != nil {
		return m.updateWorkExperienceFunc(ctx, exp)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteWorkExperience(ctx context.Context, id int64) error {
	if m.deleteWorkExperienceFunc != nil {
		return m.deleteWorkExperienceFunc(ctx, id)
	}
	return errors.New("not implemented")
}

// Certification implementations
func (m *mockRepository) GetAllCertifications(ctx context.Context) ([]models.Certification, error) {
	if m.getAllCertificationsFunc != nil {
		return m.getAllCertificationsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetCertificationByID(ctx context.Context, id int64) (*models.Certification, error) {
	if m.getCertificationByIDFunc != nil {
		return m.getCertificationByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) CreateCertification(ctx context.Context, cert *models.Certification) error {
	if m.createCertificationFunc != nil {
		return m.createCertificationFunc(ctx, cert)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateCertification(ctx context.Context, cert *models.Certification) error {
	if m.updateCertificationFunc != nil {
		return m.updateCertificationFunc(ctx, cert)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteCertification(ctx context.Context, id int64) error {
	if m.deleteCertificationFunc != nil {
		return m.deleteCertificationFunc(ctx, id)
	}
	return errors.New("not implemented")
}

// Miniature Theme implementations
func (m *mockRepository) GetAllMiniatureThemes(ctx context.Context) ([]models.MiniatureTheme, error) {
	if m.getAllMiniatureThemesFunc != nil {
		return m.getAllMiniatureThemesFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetMiniatureThemeByID(ctx context.Context, id int64) (*models.MiniatureTheme, error) {
	if m.getMiniatureThemeByIDFunc != nil {
		return m.getMiniatureThemeByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) CreateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error {
	if m.createMiniatureThemeFunc != nil {
		return m.createMiniatureThemeFunc(ctx, theme)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateMiniatureTheme(ctx context.Context, theme *models.MiniatureTheme) error {
	if m.updateMiniatureThemeFunc != nil {
		return m.updateMiniatureThemeFunc(ctx, theme)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteMiniatureTheme(ctx context.Context, id int64) error {
	if m.deleteMiniatureThemeFunc != nil {
		return m.deleteMiniatureThemeFunc(ctx, id)
	}
	return errors.New("not implemented")
}

// Miniature Project implementations
func (m *mockRepository) GetAllMiniatureProjects(ctx context.Context) ([]models.MiniatureProject, error) {
	if m.getAllMiniatureProjectsFunc != nil {
		return m.getAllMiniatureProjectsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetMiniatureProjectByID(ctx context.Context, id int64) (*models.MiniatureProject, error) {
	if m.getMiniatureProjectByIDFunc != nil {
		return m.getMiniatureProjectByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) CreateMiniatureProject(ctx context.Context, project *models.MiniatureProject) error {
	if m.createMiniatureProjectFunc != nil {
		return m.createMiniatureProjectFunc(ctx, project)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateMiniatureProject(ctx context.Context, project *models.MiniatureProject) error {
	if m.updateMiniatureProjectFunc != nil {
		return m.updateMiniatureProjectFunc(ctx, project)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteMiniatureProject(ctx context.Context, id int64) error {
	if m.deleteMiniatureProjectFunc != nil {
		return m.deleteMiniatureProjectFunc(ctx, id)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) AddImageToProject(ctx context.Context, miniatureFile *models.MiniatureFile) error {
	if m.addImageToProjectFunc != nil {
		return m.addImageToProjectFunc(ctx, miniatureFile)
	}
	return errors.New("not implemented")
}

// Miniature Paint implementations
func (m *mockRepository) GetAllMiniaturePaints(ctx context.Context) ([]models.MiniaturePaint, error) {
	if m.getAllMiniaturePaintsFunc != nil {
		return m.getAllMiniaturePaintsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetMiniaturePaintByID(ctx context.Context, id int64) (*models.MiniaturePaint, error) {
	if m.getMiniaturePaintByIDFunc != nil {
		return m.getMiniaturePaintByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) CreateMiniaturePaint(ctx context.Context, paint *models.MiniaturePaint) error {
	if m.createMiniaturePaintFunc != nil {
		return m.createMiniaturePaintFunc(ctx, paint)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateMiniaturePaint(ctx context.Context, paint *models.MiniaturePaint) error {
	if m.updateMiniaturePaintFunc != nil {
		return m.updateMiniaturePaintFunc(ctx, paint)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteMiniaturePaint(ctx context.Context, id int64) error {
	if m.deleteMiniaturePaintFunc != nil {
		return m.deleteMiniaturePaintFunc(ctx, id)
	}
	return errors.New("not implemented")
}

// Skill implementations
func (m *mockRepository) GetAllSkills(ctx context.Context) ([]models.Skill, error) {
	if m.getAllSkillsFunc != nil {
		return m.getAllSkillsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetSkillByID(ctx context.Context, id int64) (*models.Skill, error) {
	if m.getSkillByIDFunc != nil {
		return m.getSkillByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) CreateSkill(ctx context.Context, skill *models.Skill) error {
	if m.createSkillFunc != nil {
		return m.createSkillFunc(ctx, skill)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateSkill(ctx context.Context, skill *models.Skill) error {
	if m.updateSkillFunc != nil {
		return m.updateSkillFunc(ctx, skill)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteSkill(ctx context.Context, id int64) error {
	if m.deleteSkillFunc != nil {
		return m.deleteSkillFunc(ctx, id)
	}
	return errors.New("not implemented")
}

// Skill Type implementations
func (m *mockRepository) GetAllSkillTypes(ctx context.Context) ([]models.SkillType, error) {
	if m.getAllSkillTypesFunc != nil {
		return m.getAllSkillTypesFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetSkillTypeByID(ctx context.Context, id int64) (*models.SkillType, error) {
	if m.getSkillTypeByIDFunc != nil {
		return m.getSkillTypeByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) CreateSkillType(ctx context.Context, skillType *models.SkillType) error {
	if m.createSkillTypeFunc != nil {
		return m.createSkillTypeFunc(ctx, skillType)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdateSkillType(ctx context.Context, skillType *models.SkillType) error {
	if m.updateSkillTypeFunc != nil {
		return m.updateSkillTypeFunc(ctx, skillType)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeleteSkillType(ctx context.Context, id int64) error {
	if m.deleteSkillTypeFunc != nil {
		return m.deleteSkillTypeFunc(ctx, id)
	}
	return errors.New("not implemented")
}

// Portfolio Project implementations
func (m *mockRepository) GetAllPortfolioProjects(ctx context.Context) ([]models.PortfolioProject, error) {
	if m.getAllPortfolioProjectsFunc != nil {
		return m.getAllPortfolioProjectsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) GetPortfolioProjectByID(ctx context.Context, id int64) (*models.PortfolioProject, error) {
	if m.getPortfolioProjectByIDFunc != nil {
		return m.getPortfolioProjectByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *mockRepository) CreatePortfolioProject(ctx context.Context, project *models.PortfolioProject) error {
	if m.createPortfolioProjectFunc != nil {
		return m.createPortfolioProjectFunc(ctx, project)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) UpdatePortfolioProject(ctx context.Context, project *models.PortfolioProject) error {
	if m.updatePortfolioProjectFunc != nil {
		return m.updatePortfolioProjectFunc(ctx, project)
	}
	return errors.New("not implemented")
}

func (m *mockRepository) DeletePortfolioProject(ctx context.Context, id int64) error {
	if m.deletePortfolioProjectFunc != nil {
		return m.deletePortfolioProjectFunc(ctx, id)
	}
	return errors.New("not implemented")
}

// Image implementations
func (m *mockRepository) DeleteImage(ctx context.Context, id int64) error {
	if m.deleteImageFunc != nil {
		return m.deleteImageFunc(ctx, id)
	}
	return errors.New("not implemented")
}

// =============================================================================
// Test Helpers
// =============================================================================

func setupTestHandler(t *testing.T) (*Handler, *mockRepository) {
	t.Helper()

	gin.SetMode(gin.TestMode)
	mockRepo := &mockRepository{}
	handler := New(mockRepo)

	return handler, mockRepo
}

func setupTestRouter(t *testing.T, handler *Handler) *gin.Engine {
	t.Helper()

	router := gin.New()
	return router
}

func createTestCertification() models.Certification {
	return models.Certification{
		ID:            1,
		Name:          testCertName,
		Issuer:        testCertIssuer,
		IssueDate:     testCertIssueDate,
		CredentialID:  testCertBadgeID,
		CredentialURL: testCertBadgeURL,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func createTestSkill() models.Skill {
	return models.Skill{
		ID:           1,
		Skill:        testSkillName,
		SkillTypeID:  1,
		IsVisible:    true,
		DisplayOrder: 1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func createTestSkillType() models.SkillType {
	return models.SkillType{
		ID:           1,
		Name:         testSkillTypeName,
		Description:  "Languages for software development",
		DisplayOrder: 1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func createTestWorkExperience() models.WorkExperience {
	return models.WorkExperience{
		ID:          1,
		Company:     testCompanyName,
		Position:    testPosition,
		Description: "Building awesome software",
		StartDate:   testStartDate,
		IsCurrent:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func performRequest(router *gin.Engine, method, path string, body interface{}) *httptest.ResponseRecorder {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}

	req, _ := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// =============================================================================
// Constructor Tests
// =============================================================================

func TestNewHandler(t *testing.T) {
	mockRepo := &mockRepository{}
	handler := New(mockRepo)

	if handler == nil {
		t.Error("New() should return non-nil handler")
	}
}

// =============================================================================
// Certification Handler Tests
// =============================================================================

func TestGetAllCertifications_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/certifications", handler.GetAllCertifications)

	expectedCerts := []models.Certification{createTestCertification()}
	mockRepo.getAllCertificationsFunc = func(ctx context.Context) ([]models.Certification, error) {
		return expectedCerts, nil
	}

	w := performRequest(router, "GET", "/certifications", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetAllCertifications() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.Certification
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GetAllCertifications() returned %d items, want 1", len(result))
	}

	if result[0].Name != testCertName {
		t.Errorf("GetAllCertifications() name = %s, want %s", result[0].Name, testCertName)
	}
}

func TestGetAllCertifications_Empty(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/certifications", handler.GetAllCertifications)

	mockRepo.getAllCertificationsFunc = func(ctx context.Context) ([]models.Certification, error) {
		return []models.Certification{}, nil
	}

	w := performRequest(router, "GET", "/certifications", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetAllCertifications() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.Certification
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("GetAllCertifications() returned %d items, want 0", len(result))
	}
}

func TestGetAllCertifications_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/certifications", handler.GetAllCertifications)

	mockRepo.getAllCertificationsFunc = func(ctx context.Context) ([]models.Certification, error) {
		return nil, errors.New("database connection failed")
	}

	w := performRequest(router, "GET", "/certifications", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetAllCertifications() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

func TestGetCertificationByID_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/certifications/:id", handler.GetCertificationByID)

	expectedCert := createTestCertification()
	mockRepo.getCertificationByIDFunc = func(ctx context.Context, id int64) (*models.Certification, error) {
		if id != 1 {
			return nil, gorm.ErrRecordNotFound
		}
		return &expectedCert, nil
	}

	w := performRequest(router, "GET", "/certifications/1", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetCertificationByID() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result models.Certification
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result.Name != testCertName {
		t.Errorf("GetCertificationByID() name = %s, want %s", result.Name, testCertName)
	}
}

func TestGetCertificationByID_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/certifications/:id", handler.GetCertificationByID)

	mockRepo.getCertificationByIDFunc = func(ctx context.Context, id int64) (*models.Certification, error) {
		return nil, gorm.ErrRecordNotFound
	}

	w := performRequest(router, "GET", "/certifications/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("GetCertificationByID() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestGetCertificationByID_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/certifications/:id", handler.GetCertificationByID)

	w := performRequest(router, "GET", "/certifications/invalid", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("GetCertificationByID() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestCreateCertification_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.POST("/certifications", handler.CreateCertification)

	mockRepo.createCertificationFunc = func(ctx context.Context, cert *models.Certification) error {
		cert.ID = 1
		return nil
	}

	newCert := map[string]interface{}{
		"name":      testCertName,
		"issuer":    testCertIssuer,
		"issueDate": testCertIssueDate,
	}

	w := performRequest(router, "POST", "/certifications", newCert)

	if w.Code != http.StatusCreated {
		t.Errorf("CreateCertification() status = %d, want %d", w.Code, http.StatusCreated)
	}

	location := w.Header().Get("Location")
	if location != "/certifications/1" {
		t.Errorf("CreateCertification() Location = %s, want /certifications/1", location)
	}
}

func TestCreateCertification_ValidationError(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.POST("/certifications", handler.CreateCertification)

	// Missing required fields
	invalidCert := map[string]interface{}{
		"name": testCertName,
		// missing issuer and issueDate
	}

	w := performRequest(router, "POST", "/certifications", invalidCert)

	if w.Code != http.StatusBadRequest {
		t.Errorf("CreateCertification() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestCreateCertification_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.POST("/certifications", handler.CreateCertification)

	mockRepo.createCertificationFunc = func(ctx context.Context, cert *models.Certification) error {
		return errors.New("database error")
	}

	newCert := map[string]interface{}{
		"name":      testCertName,
		"issuer":    testCertIssuer,
		"issueDate": testCertIssueDate,
	}

	w := performRequest(router, "POST", "/certifications", newCert)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("CreateCertification() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

func TestUpdateCertification_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.PUT("/certifications/:id", handler.UpdateCertification)

	mockRepo.updateCertificationFunc = func(ctx context.Context, cert *models.Certification) error {
		return nil
	}

	updateCert := map[string]interface{}{
		"name":      "Updated Certification",
		"issuer":    testCertIssuer,
		"issueDate": testCertIssueDate,
	}

	w := performRequest(router, "PUT", "/certifications/1", updateCert)

	if w.Code != http.StatusOK {
		t.Errorf("UpdateCertification() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestUpdateCertification_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.PUT("/certifications/:id", handler.UpdateCertification)

	mockRepo.updateCertificationFunc = func(ctx context.Context, cert *models.Certification) error {
		return gorm.ErrRecordNotFound
	}

	updateCert := map[string]interface{}{
		"name":      "Updated Certification",
		"issuer":    testCertIssuer,
		"issueDate": testCertIssueDate,
	}

	w := performRequest(router, "PUT", "/certifications/999", updateCert)

	if w.Code != http.StatusNotFound {
		t.Errorf("UpdateCertification() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestUpdateCertification_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.PUT("/certifications/:id", handler.UpdateCertification)

	w := performRequest(router, "PUT", "/certifications/invalid", map[string]interface{}{})

	if w.Code != http.StatusBadRequest {
		t.Errorf("UpdateCertification() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestDeleteCertification_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.DELETE("/certifications/:id", handler.DeleteCertification)

	mockRepo.deleteCertificationFunc = func(ctx context.Context, id int64) error {
		return nil
	}

	w := performRequest(router, "DELETE", "/certifications/1", nil)

	if w.Code != http.StatusNoContent {
		t.Errorf("DeleteCertification() status = %d, want %d", w.Code, http.StatusNoContent)
	}
}

func TestDeleteCertification_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.DELETE("/certifications/:id", handler.DeleteCertification)

	mockRepo.deleteCertificationFunc = func(ctx context.Context, id int64) error {
		return gorm.ErrRecordNotFound
	}

	w := performRequest(router, "DELETE", "/certifications/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("DeleteCertification() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestDeleteCertification_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.DELETE("/certifications/:id", handler.DeleteCertification)

	w := performRequest(router, "DELETE", "/certifications/invalid", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("DeleteCertification() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

// =============================================================================
// Skill Handler Tests
// =============================================================================

func TestGetAllSkills_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/skills", handler.GetAllSkills)

	expectedSkills := []models.Skill{createTestSkill()}
	mockRepo.getAllSkillsFunc = func(ctx context.Context) ([]models.Skill, error) {
		return expectedSkills, nil
	}

	w := performRequest(router, "GET", "/skills", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetAllSkills() status = %d, want %d", w.Code, http.StatusOK)
	}

	var result []models.Skill
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GetAllSkills() returned %d items, want 1", len(result))
	}
}

func TestGetAllSkills_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/skills", handler.GetAllSkills)

	mockRepo.getAllSkillsFunc = func(ctx context.Context) ([]models.Skill, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(router, "GET", "/skills", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetAllSkills() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

func TestGetSkillByID_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/skills/:id", handler.GetSkillByID)

	expectedSkill := createTestSkill()
	mockRepo.getSkillByIDFunc = func(ctx context.Context, id int64) (*models.Skill, error) {
		return &expectedSkill, nil
	}

	w := performRequest(router, "GET", "/skills/1", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetSkillByID() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetSkillByID_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/skills/:id", handler.GetSkillByID)

	mockRepo.getSkillByIDFunc = func(ctx context.Context, id int64) (*models.Skill, error) {
		return nil, gorm.ErrRecordNotFound
	}

	w := performRequest(router, "GET", "/skills/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("GetSkillByID() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestGetSkillByID_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/skills/:id", handler.GetSkillByID)

	w := performRequest(router, "GET", "/skills/invalid", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("GetSkillByID() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestCreateSkill_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.POST("/skills", handler.CreateSkill)

	mockRepo.createSkillFunc = func(ctx context.Context, skill *models.Skill) error {
		skill.ID = 1
		return nil
	}

	newSkill := map[string]interface{}{
		"skill":       testSkillName,
		"skillTypeId": 1,
	}

	w := performRequest(router, "POST", "/skills", newSkill)

	if w.Code != http.StatusCreated {
		t.Errorf("CreateSkill() status = %d, want %d", w.Code, http.StatusCreated)
	}

	location := w.Header().Get("Location")
	if location != "/skills/1" {
		t.Errorf("CreateSkill() Location = %s, want /skills/1", location)
	}
}

func TestCreateSkill_ValidationError(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.POST("/skills", handler.CreateSkill)

	// Missing required fields
	invalidSkill := map[string]interface{}{
		"skill": testSkillName,
		// missing skillTypeId
	}

	w := performRequest(router, "POST", "/skills", invalidSkill)

	if w.Code != http.StatusBadRequest {
		t.Errorf("CreateSkill() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestUpdateSkill_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.PUT("/skills/:id", handler.UpdateSkill)

	mockRepo.updateSkillFunc = func(ctx context.Context, skill *models.Skill) error {
		return nil
	}

	updateSkill := map[string]interface{}{
		"skill":       "Updated Skill",
		"skillTypeId": 1,
	}

	w := performRequest(router, "PUT", "/skills/1", updateSkill)

	if w.Code != http.StatusOK {
		t.Errorf("UpdateSkill() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestUpdateSkill_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.PUT("/skills/:id", handler.UpdateSkill)

	mockRepo.updateSkillFunc = func(ctx context.Context, skill *models.Skill) error {
		return gorm.ErrRecordNotFound
	}

	updateSkill := map[string]interface{}{
		"skill":       "Updated Skill",
		"skillTypeId": 1,
	}

	w := performRequest(router, "PUT", "/skills/999", updateSkill)

	if w.Code != http.StatusNotFound {
		t.Errorf("UpdateSkill() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestDeleteSkill_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.DELETE("/skills/:id", handler.DeleteSkill)

	mockRepo.deleteSkillFunc = func(ctx context.Context, id int64) error {
		return nil
	}

	w := performRequest(router, "DELETE", "/skills/1", nil)

	if w.Code != http.StatusNoContent {
		t.Errorf("DeleteSkill() status = %d, want %d", w.Code, http.StatusNoContent)
	}
}

func TestDeleteSkill_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.DELETE("/skills/:id", handler.DeleteSkill)

	mockRepo.deleteSkillFunc = func(ctx context.Context, id int64) error {
		return gorm.ErrRecordNotFound
	}

	w := performRequest(router, "DELETE", "/skills/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("DeleteSkill() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

// =============================================================================
// Skill Type Handler Tests
// =============================================================================

func TestGetAllSkillTypes_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/skill-types", handler.GetAllSkillTypes)

	expectedSkillTypes := []models.SkillType{createTestSkillType()}
	mockRepo.getAllSkillTypesFunc = func(ctx context.Context) ([]models.SkillType, error) {
		return expectedSkillTypes, nil
	}

	w := performRequest(router, "GET", "/skill-types", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetAllSkillTypes() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetSkillTypeByID_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/skill-types/:id", handler.GetSkillTypeByID)

	expectedSkillType := createTestSkillType()
	mockRepo.getSkillTypeByIDFunc = func(ctx context.Context, id int64) (*models.SkillType, error) {
		return &expectedSkillType, nil
	}

	w := performRequest(router, "GET", "/skill-types/1", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetSkillTypeByID() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetSkillTypeByID_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/skill-types/:id", handler.GetSkillTypeByID)

	mockRepo.getSkillTypeByIDFunc = func(ctx context.Context, id int64) (*models.SkillType, error) {
		return nil, gorm.ErrRecordNotFound
	}

	w := performRequest(router, "GET", "/skill-types/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("GetSkillTypeByID() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestCreateSkillType_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.POST("/skill-types", handler.CreateSkillType)

	mockRepo.createSkillTypeFunc = func(ctx context.Context, skillType *models.SkillType) error {
		skillType.ID = 1
		return nil
	}

	newSkillType := map[string]interface{}{
		"name": testSkillTypeName,
	}

	w := performRequest(router, "POST", "/skill-types", newSkillType)

	if w.Code != http.StatusCreated {
		t.Errorf("CreateSkillType() status = %d, want %d", w.Code, http.StatusCreated)
	}
}

func TestUpdateSkillType_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.PUT("/skill-types/:id", handler.UpdateSkillType)

	mockRepo.updateSkillTypeFunc = func(ctx context.Context, skillType *models.SkillType) error {
		return nil
	}

	updateSkillType := map[string]interface{}{
		"name": "Updated Skill Type",
	}

	w := performRequest(router, "PUT", "/skill-types/1", updateSkillType)

	if w.Code != http.StatusOK {
		t.Errorf("UpdateSkillType() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestDeleteSkillType_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.DELETE("/skill-types/:id", handler.DeleteSkillType)

	mockRepo.deleteSkillTypeFunc = func(ctx context.Context, id int64) error {
		return nil
	}

	w := performRequest(router, "DELETE", "/skill-types/1", nil)

	if w.Code != http.StatusNoContent {
		t.Errorf("DeleteSkillType() status = %d, want %d", w.Code, http.StatusNoContent)
	}
}

// =============================================================================
// Work Experience Handler Tests
// =============================================================================

func TestGetAllWorkExperience_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/work-experience", handler.GetAllWorkExperience)

	expectedExps := []models.WorkExperience{createTestWorkExperience()}
	mockRepo.getAllWorkExperienceFunc = func(ctx context.Context) ([]models.WorkExperience, error) {
		return expectedExps, nil
	}

	w := performRequest(router, "GET", "/work-experience", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetAllWorkExperience() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetAllWorkExperience_RepositoryError(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/work-experience", handler.GetAllWorkExperience)

	mockRepo.getAllWorkExperienceFunc = func(ctx context.Context) ([]models.WorkExperience, error) {
		return nil, errors.New("database error")
	}

	w := performRequest(router, "GET", "/work-experience", nil)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("GetAllWorkExperience() status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

func TestGetWorkExperienceByID_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/work-experience/:id", handler.GetWorkExperienceByID)

	expectedExp := createTestWorkExperience()
	mockRepo.getWorkExperienceByIDFunc = func(ctx context.Context, id int64) (*models.WorkExperience, error) {
		return &expectedExp, nil
	}

	w := performRequest(router, "GET", "/work-experience/1", nil)

	if w.Code != http.StatusOK {
		t.Errorf("GetWorkExperienceByID() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetWorkExperienceByID_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/work-experience/:id", handler.GetWorkExperienceByID)

	mockRepo.getWorkExperienceByIDFunc = func(ctx context.Context, id int64) (*models.WorkExperience, error) {
		return nil, gorm.ErrRecordNotFound
	}

	w := performRequest(router, "GET", "/work-experience/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("GetWorkExperienceByID() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestGetWorkExperienceByID_InvalidID(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/work-experience/:id", handler.GetWorkExperienceByID)

	w := performRequest(router, "GET", "/work-experience/invalid", nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("GetWorkExperienceByID() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestCreateWorkExperience_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.POST("/work-experience", handler.CreateWorkExperience)

	mockRepo.createWorkExperienceFunc = func(ctx context.Context, exp *models.WorkExperience) error {
		exp.ID = 1
		return nil
	}

	newExp := map[string]interface{}{
		"company":   testCompanyName,
		"position":  testPosition,
		"startDate": testStartDate,
	}

	w := performRequest(router, "POST", "/work-experience", newExp)

	if w.Code != http.StatusCreated {
		t.Errorf("CreateWorkExperience() status = %d, want %d", w.Code, http.StatusCreated)
	}

	location := w.Header().Get("Location")
	if location != "/work-experience/1" {
		t.Errorf("CreateWorkExperience() Location = %s, want /work-experience/1", location)
	}
}

func TestCreateWorkExperience_ValidationError(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.POST("/work-experience", handler.CreateWorkExperience)

	// Missing required fields
	invalidExp := map[string]interface{}{
		"company": testCompanyName,
		// missing position and startDate
	}

	w := performRequest(router, "POST", "/work-experience", invalidExp)

	if w.Code != http.StatusBadRequest {
		t.Errorf("CreateWorkExperience() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestUpdateWorkExperience_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.PUT("/work-experience/:id", handler.UpdateWorkExperience)

	mockRepo.updateWorkExperienceFunc = func(ctx context.Context, exp *models.WorkExperience) error {
		return nil
	}

	updateExp := map[string]interface{}{
		"company":   "Updated Company",
		"position":  testPosition,
		"startDate": testStartDate,
	}

	w := performRequest(router, "PUT", "/work-experience/1", updateExp)

	if w.Code != http.StatusOK {
		t.Errorf("UpdateWorkExperience() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestUpdateWorkExperience_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.PUT("/work-experience/:id", handler.UpdateWorkExperience)

	mockRepo.updateWorkExperienceFunc = func(ctx context.Context, exp *models.WorkExperience) error {
		return gorm.ErrRecordNotFound
	}

	updateExp := map[string]interface{}{
		"company":   "Updated Company",
		"position":  testPosition,
		"startDate": testStartDate,
	}

	w := performRequest(router, "PUT", "/work-experience/999", updateExp)

	if w.Code != http.StatusNotFound {
		t.Errorf("UpdateWorkExperience() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

func TestDeleteWorkExperience_Success(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.DELETE("/work-experience/:id", handler.DeleteWorkExperience)

	mockRepo.deleteWorkExperienceFunc = func(ctx context.Context, id int64) error {
		return nil
	}

	w := performRequest(router, "DELETE", "/work-experience/1", nil)

	if w.Code != http.StatusNoContent {
		t.Errorf("DeleteWorkExperience() status = %d, want %d", w.Code, http.StatusNoContent)
	}
}

func TestDeleteWorkExperience_NotFound(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.DELETE("/work-experience/:id", handler.DeleteWorkExperience)

	mockRepo.deleteWorkExperienceFunc = func(ctx context.Context, id int64) error {
		return gorm.ErrRecordNotFound
	}

	w := performRequest(router, "DELETE", "/work-experience/999", nil)

	if w.Code != http.StatusNotFound {
		t.Errorf("DeleteWorkExperience() status = %d, want %d", w.Code, http.StatusNotFound)
	}
}

// =============================================================================
// Context Propagation Tests
// =============================================================================

func TestContextPropagation(t *testing.T) {
	handler, mockRepo := setupTestHandler(t)
	router := setupTestRouter(t, handler)
	router.GET("/certifications", handler.GetAllCertifications)

	var receivedCtx context.Context
	mockRepo.getAllCertificationsFunc = func(ctx context.Context) ([]models.Certification, error) {
		receivedCtx = ctx
		return []models.Certification{}, nil
	}

	w := performRequest(router, "GET", "/certifications", nil)

	if w.Code != http.StatusOK {
		t.Errorf("Request failed with status %d", w.Code)
	}

	if receivedCtx == nil {
		t.Error("Context was not propagated to repository")
	}
}

// =============================================================================
// Table-Driven Tests for ID Validation
// =============================================================================

func TestInvalidIDFormats(t *testing.T) {
	handler, _ := setupTestHandler(t)
	router := setupTestRouter(t, handler)

	router.GET("/certifications/:id", handler.GetCertificationByID)
	router.GET("/skills/:id", handler.GetSkillByID)
	router.GET("/work-experience/:id", handler.GetWorkExperienceByID)

	// Note: Negative IDs are parseable by strconv.ParseInt, so they pass validation
	// and get a "not found" from the repository. Only non-numeric strings fail parsing.
	tests := []struct {
		name      string
		path      string
		invalidID string
	}{
		{"certification with string ID", "/certifications/", "abc"},
		{"certification with float ID", "/certifications/", "1.5"},
		{"skill with string ID", "/skills/", "xyz"},
		{"work experience with string ID", "/work-experience/", "test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := performRequest(router, "GET", tt.path+tt.invalidID, nil)

			if w.Code != http.StatusBadRequest {
				t.Errorf("%s: status = %d, want %d", tt.name, w.Code, http.StatusBadRequest)
			}
		})
	}
}
