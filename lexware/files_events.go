package lexware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/rasche-thalhofer/lexware-go/types"
)

type filesClient struct{ client *Client }

func (c *filesClient) Upload(ctx context.Context, filename string, content io.Reader, fileType types.FileUploadType) (*types.FileUploadResponse, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, content); err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}
	if err := writer.WriteField("type", string(fileType)); err != nil {
		return nil, fmt.Errorf("failed to write type field: %w", err)
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.client.baseURL+"/v1/files", &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.client.apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}
	var result types.FileUploadResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *filesClient) Download(ctx context.Context, id string) (io.ReadCloser, error) {
	resp, err := c.client.doRequestRaw(ctx, "GET", "/v1/files/"+id, nil, map[string]string{"Accept": "application/octet-stream"})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, &APIError{StatusCode: resp.StatusCode, Body: string(body)}
	}
	return resp.Body, nil
}

type eventSubscriptionsClient struct{ client *Client }

func (c *eventSubscriptionsClient) Create(ctx context.Context, subscription *types.EventSubscriptionCreateRequest) (*types.ActionResult, error) {
	body, err := c.client.doRequest(ctx, "POST", "/v1/event-subscriptions", subscription)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *eventSubscriptionsClient) Get(ctx context.Context, id string) (*types.EventSubscription, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/event-subscriptions/"+id, nil)
	if err != nil {
		return nil, err
	}
	var subscription types.EventSubscription
	if err := json.Unmarshal(body, &subscription); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &subscription, nil
}

func (c *eventSubscriptionsClient) List(ctx context.Context) ([]types.EventSubscription, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/event-subscriptions", nil)
	if err != nil {
		return nil, err
	}
	var response struct {
		Content []types.EventSubscription `json:"content"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return response.Content, nil
}

func (c *eventSubscriptionsClient) Delete(ctx context.Context, id string) error {
	_, err := c.client.doRequest(ctx, "DELETE", "/v1/event-subscriptions/"+id, nil)
	return err
}
