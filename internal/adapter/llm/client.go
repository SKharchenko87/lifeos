package llm

import (
	"context"
	"encoding/json"
	"lifeos/internal/domain"
	"net/http"

	"github.com/google/uuid"
)

type chatRequest struct {
	Model          string         `json:"model"`
	Messages       []chatMessage  `json:"messages"`
	Temperature    float64        `json:"temperature"`
	ResponseFormat responseFormat `json:"response_format"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type responseFormat struct {
	Format string `json:"type"`
}

type chatResponse struct {
	Choices []chatChoice `json:"choices"`
}

type chatChoice struct {
	Message responseMessage `json:"message"`
}

type responseMessage struct {
	Content string `json:"content"`
}

type content struct {
	Items []item `json:"items"`
}

type item struct {
	Name         string  `json:"name"`
	Weight       int     `json:"weight"`
	KCal         int     `json:"kcal"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Carbohydrate float64 `json:"carbohydrate"`
}

type Parser interface {
	Parse(ctx context.Context, text string) ([]domain.MealItem, error)
}

type Client struct {
	baseUrl    string
	model      string
	httpClient *http.Client
}

func NewClient(baseUrl string, model string, httpClient *http.Client) *Client {
	return &Client{baseUrl: baseUrl, model: model, httpClient: httpClient}
}

func (c Client) Parse(ctx context.Context, text string) ([]domain.MealItem, error) {
	chatResponse := chatResponse{}
	err := json.Unmarshal([]byte(text), &chatResponse)
	if err != nil {
		return nil, err
	}
	result := []domain.MealItem{}
	for _, choice := range chatResponse.Choices {
		content := choice.Message.Content
		item := item{}
		err = json.Unmarshal([]byte(content), &item)
		if err != nil {
			return nil, err
		}
		meal := domain.MealItem{
			ID:           uuid.New(),
			Name:         item.Name,
			Weight:       item.Weight,
			KCal:         item.KCal,
			Protein:      item.Protein,
			Fat:          item.Fat,
			Carbohydrate: item.Carbohydrate,
		}
		result = append(result, meal)
	}

	return result, nil
}
