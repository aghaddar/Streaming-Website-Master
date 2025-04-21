package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ConsumetService struct {
	BaseURL string
}

func NewConsumetService(baseURL string) *ConsumetService {
	return &ConsumetService{BaseURL: baseURL}
}

func (s *ConsumetService) SearchAnime(provider, searchTerm string, page int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/anime/%s/%s?page=%d", s.BaseURL, provider, searchTerm, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ConsumetService) GetAnimeInfo(provider, seriesID string, episodePage int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/anime/%s/info/%s?episodePage=%d", s.BaseURL, provider, seriesID, episodePage)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ConsumetService) WatchEpisode(provider, episodeID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/anime/%s/watch?episodeId=%s", s.BaseURL, provider, episodeID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
