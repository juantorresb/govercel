package dtos

import (
	echo "github.com/tbxark/g4vercel"
)

type MediaDto struct {
	StatusCode int     `json:"statusCode"`
	Message    string  `json:"message"`
	Data       *echo.H `json:"data"`
}
