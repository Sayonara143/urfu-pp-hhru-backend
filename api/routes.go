package api

func (s *Server) setupRoutes() {
	// Группа маршрутов для аутентификации
	auth := s.e.Group("/auth")
	auth.POST("/register", s.register)
	auth.POST("/login", s.login)

	// Группа маршрутов для пользователей
	users := s.e.Group("/users")
	users.GET("/:id", s.getUser)
	users.GET("", s.getAllUsers)
	users.PUT("/:id", s.updateUser)
	users.DELETE("/:id", s.deleteUser)

	// Группа маршрутов для профилей
	profiles := s.e.Group("/profiles")
	profiles.GET("/student/:id", s.getStudentProfile)
	profiles.GET("/employer/:id", s.getEmployerProfile)
	profiles.PUT("/student/:id", s.updateStudentProfile)
	profiles.PUT("/employer/:id", s.updateEmployerProfile)

	// Группа маршрутов для вакансий
	vacancies := s.e.Group("/vacancies")
	vacancies.GET("/:id", s.getVacancy)
	vacancies.GET("", s.getAllVacancies)
	vacancies.POST("", s.createVacancy)
	vacancies.PUT("/:id", s.updateVacancy)
	vacancies.DELETE("/:id", s.deleteVacancy)

	// Группа маршрутов для заявок
	applications := s.e.Group("/applications")
	applications.GET("/:id", s.getApplication)
	applications.GET("", s.getAllApplications)
	applications.POST("", s.createApplication)
	applications.PUT("/:id", s.updateApplication)
	applications.DELETE("/:id", s.deleteApplication)

	// Группа маршрутов для собеседований
	interviews := s.e.Group("/interviews")
	interviews.GET("/:id", s.getInterview)
	interviews.GET("", s.getAllInterviews)
	interviews.POST("", s.createInterview)
	interviews.PUT("/:id", s.updateInterview)
	interviews.DELETE("/:id", s.deleteInterview)

	// Группа маршрутов для отзывов
	reviews := s.e.Group("/reviews")
	reviews.GET("/:id", s.getReview)
	reviews.GET("", s.getAllReviews)
	reviews.POST("", s.createReview)
	reviews.PUT("/:id", s.updateReview)
	reviews.DELETE("/:id", s.deleteReview)

	// Группа маршрутов для чёрного списка
	blacklist := s.e.Group("/blacklist")
	blacklist.GET("/:id", s.getBlacklistEntry)
	blacklist.GET("", s.getAllBlacklistEntries)
	blacklist.POST("", s.addBlacklistEntry)
	blacklist.DELETE("/:id", s.removeBlacklistEntry)

	// Группа маршрутов для резюме
	resumes := s.e.Group("/resumes")
	resumes.GET("/:id", s.getResume)
	resumes.GET("", s.getAllResumes)
	resumes.POST("", s.createResume)
	resumes.PUT("/:id", s.updateResume)
	resumes.DELETE("/:id", s.deleteResume)

	// Группа маршрутов для поиска и сохранённых запросов
	search := s.e.Group("/search")
	search.GET("/logs", s.getSearchLogs)
	search.GET("/saved", s.getSavedSearches)
	search.POST("/saved", s.saveSearchQuery)
	search.DELETE("/saved/:id", s.deleteSavedSearchQuery)
}
