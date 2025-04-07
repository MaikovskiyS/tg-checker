package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type telegramClient struct {
	httpClient *http.Client
}

func NewTelegramClient(timeout time.Duration) *telegramClient {
	return &telegramClient{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// GetChatMember проверяет, является ли пользователь участником канала
func (c *telegramClient) GetChatMember(botToken, channelLink string, userID int64) (bool, error) {
	chatID, err := c.extractChatID(channelLink)
	if err != nil {
		return false, err
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/getChatMember", botToken)

	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("user_id", fmt.Sprintf("%d", userID))

	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return false, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result struct {
		OK     bool `json:"ok"`
		Result struct {
			Status string `json:"status"`
		} `json:"result"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if !result.OK {
		return false, fmt.Errorf("ошибка Telegram API: %s", result.Description)
	}
	return result.Result.Status == "member" ||
		result.Result.Status == "administrator" ||
		result.Result.Status == "creator", nil
}

// extractChatID извлекает ID канала из ссылки
func (c *telegramClient) extractChatID(channelLink string) (string, error) {
	if strings.HasPrefix(channelLink, "@") {
		return channelLink, nil
	}

	if strings.HasPrefix(channelLink, "-100") {
		return channelLink, nil
	}

	parsedURL, err := url.Parse(channelLink)
	if err != nil {
		return "", err
	}

	path := parsedURL.Path
	if path == "" {
		return "", fmt.Errorf("некорректная ссылка на канал: %s", channelLink)
	}

	path = strings.TrimPrefix(path, "/")

	return "@" + path, nil
}
