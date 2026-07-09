package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"test/3-struct/bins"
	"time"
)

type CloudStorage struct {
	baseUrl string
	KEY     string
}

type Storage interface {
	SaveBin(bins.Bin) error
	GetBin(string) (bins.Bin, error)
	// UpdatedBin(bins.Bin) error
	DeleteBin(string) error
	GetAllBins() (bins.BinList, error)
}

func (s *CloudStorage) SaveBin(bin bins.Bin) error {
	data, err := json.Marshal(bin)
	if err != nil {
		return err
	}

	url, _ := url.JoinPath(s.baseUrl, "/b")
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	fmt.Println("req", req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", s.KEY)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to save bin: status %d, body: %s", resp.StatusCode, string(body))
	}
	return nil
}

func (s *CloudStorage) GetBin(id string) (bins.Bin, error) {
	// 1. Формируем URL
	url, err := url.JoinPath(s.baseUrl, "/b/", id)
	if err != nil {
		return bins.Bin{}, err // ← возвращаем пустой Bin, не nil
	}

	// 2. Создаем HTTP запрос (не http.Get, а http.NewRequest)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return bins.Bin{}, err
	}

	// 3. Добавляем заголовок авторизации
	req.Header.Set("X-Master-Key", s.KEY)

	// 4. Отправляем запрос
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return bins.Bin{}, err
	}
	defer resp.Body.Close() // ← ОБЯЗАТЕЛЬНО закрываем body!

	// 5. Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return bins.Bin{}, fmt.Errorf("failed to get bin: status %d, body: %s", resp.StatusCode, string(body))
	}

	// 6. Парсим JSON из ответа в структуру bins.Bin
	var bin bins.Bin
	err = json.NewDecoder(resp.Body).Decode(&bin)
	if err != nil {
		return bins.Bin{}, err
	}
	fmt.Println("bin", bin, resp.Status)
	// 7. Возвращаем bin
	return bin, nil
}

// func (s *CloudStorage) UpdateBin(bin bins.Bin) error {}

func (s *CloudStorage) DeleteBin(id string) error {
	url, err := url.JoinPath(s.baseUrl, "/b/", id)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Master-Key", s.KEY)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil

}

func (s *CloudStorage) GetAllBins() (bins.BinList, error) {
	url, err := url.JoinPath(s.baseUrl, "/l")
	if err != nil {
		return bins.BinList{}, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return bins.BinList{}, err
	}
	req.Header.Set("X-Master-Key", s.KEY)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return bins.BinList{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return bins.BinList{}, fmt.Errorf("failed to get all bins: status %d, body: %s", resp.StatusCode, string(body))
	}

	// 6. Промежуточная структура для парсинга
	type jsonbinResponse struct {
		Record []bins.Bin `json:"record"`
	}

	var response jsonbinResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return bins.BinList{}, err
	}

	// 7. ВАЖНО: Проверяем, что Record не nil!
	if response.Record == nil {
		return bins.BinList{}, nil // возвращаем пустой список
	}

	// 8. Возвращаем BinList
	return bins.BinList{Bin: response.Record}, nil
}

func NewStorage(baseUrl string, KEY string) *CloudStorage {
	return &CloudStorage{
		baseUrl: baseUrl,
		KEY:     KEY,
	}
}
