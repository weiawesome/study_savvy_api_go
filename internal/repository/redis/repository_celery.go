package redis

import (
	"context"
	"github.com/google/uuid"
	"study_savvy_api_go/api/model"
	"study_savvy_api_go/api/utils"
)

func (r *Repository) OcrMission(id string, file string, prompt string, key model.ApiKey, token model.AccessToken) error {
	body := utils.CeleryBody{
		Position: []utils.CeleryBodyPosition{},
		Args: utils.CeleryOcrBodyArgs{
			Id:             id,
			File:           file,
			Prompt:         prompt,
			ApiKey:         key.Key,
			KeyApiKey:      key.AesKey,
			AccessToken:    token.Token,
			KeyAccessToken: token.AesKey,
		},
	}
	bodyBase64, err := body.ToJsonBase64()
	if err != nil {
		return err
	}

	headers := utils.CeleryHeaders{
		Id:           id,
		Task:         utils.EnvCeleryTask() + ".OCR_predict",
		Lang:         "Go",
		RootId:       id,
		IgnoreResult: false,
		Retries:      0,
		Origin:       "Go-User",
	}

	properties := utils.CeleryProperties{
		Priority:      0,
		CorrelationId: id,
		ReplyTo:       id,
		DeliveryMode:  2,
		DeliveryInfo:  utils.CeleryPropertyDeliverInfo{RoutingKey: utils.CeleryRedisKey, Exchange: ""},
		BodyEncoding:  utils.BodyEncoding,
		DeliveryTag:   id,
	}

	task := utils.CeleryTask{Body: bodyBase64, ContentEncoding: utils.ContentEncoding, ContentType: utils.ContentType, Headers: headers, Properties: properties}

	jsonData, err := task.ToJson()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := r.client.LPush(ctx, utils.CeleryRedisKey, jsonData).Err(); err != nil {
		return err
	}
	return nil
}
func (r *Repository) OcrTextMission(id string, content string, prompt string, key model.ApiKey, token model.AccessToken) error {
	body := utils.CeleryBody{
		Position: []utils.CeleryBodyPosition{},
		Args: utils.CeleryOcrTextBodyArgs{
			Id:             id,
			Prompt:         prompt,
			Content:        content,
			ApiKey:         key.Key,
			KeyApiKey:      key.AesKey,
			AccessToken:    token.Token,
			KeyAccessToken: token.AesKey,
		},
	}
	bodyBase64, err := body.ToJsonBase64()
	if err != nil {
		return err
	}

	headers := utils.CeleryHeaders{
		Id:           id,
		Task:         utils.EnvCeleryTask() + ".OCR_predict_Text",
		Lang:         "Go",
		RootId:       id,
		IgnoreResult: false,
		Retries:      0,
		Origin:       "Go-User",
	}

	properties := utils.CeleryProperties{
		Priority:      0,
		CorrelationId: id,
		ReplyTo:       id,
		DeliveryMode:  2,
		DeliveryInfo:  utils.CeleryPropertyDeliverInfo{RoutingKey: utils.CeleryRedisKey, Exchange: ""},
		BodyEncoding:  utils.BodyEncoding,
		DeliveryTag:   id,
	}

	task := utils.CeleryTask{Body: bodyBase64, ContentEncoding: utils.ContentEncoding, ContentType: utils.ContentType, Headers: headers, Properties: properties}

	jsonData, err := task.ToJson()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := r.client.LPush(ctx, utils.CeleryRedisKey, jsonData).Err(); err != nil {
		return err
	}
	return nil
}
func (r *Repository) AsrMission(id string, file string, prompt string, key model.ApiKey, token model.AccessToken) error {
	body := utils.CeleryBody{
		Position: []utils.CeleryBodyPosition{},
		Args: utils.CeleryAsrBodyArgs{
			Id:             id,
			File:           file,
			Prompt:         prompt,
			ApiKey:         key.Key,
			KeyApiKey:      key.AesKey,
			AccessToken:    token.Token,
			KeyAccessToken: token.AesKey,
		},
	}
	bodyBase64, err := body.ToJsonBase64()
	if err != nil {
		return err
	}

	headers := utils.CeleryHeaders{
		Id:           id,
		Task:         utils.EnvCeleryTask() + ".ASR_predict",
		Lang:         "Go",
		RootId:       id,
		IgnoreResult: false,
		Retries:      0,
		Origin:       "Go-User",
	}

	properties := utils.CeleryProperties{
		Priority:      0,
		CorrelationId: id,
		ReplyTo:       id,
		DeliveryMode:  2,
		DeliveryInfo:  utils.CeleryPropertyDeliverInfo{RoutingKey: utils.CeleryRedisKey, Exchange: ""},
		BodyEncoding:  utils.BodyEncoding,
		DeliveryTag:   id,
	}

	task := utils.CeleryTask{Body: bodyBase64, ContentEncoding: utils.ContentEncoding, ContentType: utils.ContentType, Headers: headers, Properties: properties}

	jsonData, err := task.ToJson()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := r.client.LPush(ctx, utils.CeleryRedisKey, jsonData).Err(); err != nil {
		return err
	}
	return nil
}
func (r *Repository) NlpEditOcrMission(id string, content string, prompt string, key model.ApiKey, token model.AccessToken) error {
	body := utils.CeleryBody{
		Position: []utils.CeleryBodyPosition{},
		Args: utils.CeleryNlpEditOcrBodyArgs{
			Id:             id,
			Prompt:         prompt,
			Content:        content,
			ApiKey:         key.Key,
			KeyApiKey:      key.AesKey,
			AccessToken:    token.Token,
			KeyAccessToken: token.AesKey,
		},
	}
	bodyBase64, err := body.ToJsonBase64()
	if err != nil {
		return err
	}

	headers := utils.CeleryHeaders{
		Id:           id,
		Task:         utils.EnvCeleryTask() + ".NLP_edit_OCR",
		Lang:         "Go",
		RootId:       id,
		IgnoreResult: false,
		Retries:      0,
		Origin:       "Go-User",
	}

	properties := utils.CeleryProperties{
		Priority:      0,
		CorrelationId: id,
		ReplyTo:       id,
		DeliveryMode:  2,
		DeliveryInfo:  utils.CeleryPropertyDeliverInfo{RoutingKey: utils.CeleryRedisKey, Exchange: ""},
		BodyEncoding:  utils.BodyEncoding,
		DeliveryTag:   id,
	}

	task := utils.CeleryTask{Body: bodyBase64, ContentEncoding: utils.ContentEncoding, ContentType: utils.ContentType, Headers: headers, Properties: properties}

	jsonData, err := task.ToJson()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := r.client.LPush(ctx, utils.CeleryRedisKey, jsonData).Err(); err != nil {
		return err
	}
	return nil
}
func (r *Repository) NlpEditAsrMission(id string, content string, prompt string, key model.ApiKey, token model.AccessToken) error {
	body := utils.CeleryBody{
		Position: []utils.CeleryBodyPosition{},
		Args: utils.CeleryNlpEditAsrBodyArgs{
			Id:             id,
			Prompt:         prompt,
			Content:        content,
			ApiKey:         key.Key,
			KeyApiKey:      key.AesKey,
			AccessToken:    token.Token,
			KeyAccessToken: token.AesKey,
		},
	}
	bodyBase64, err := body.ToJsonBase64()
	if err != nil {
		return err
	}

	headers := utils.CeleryHeaders{
		Id:           id,
		Task:         utils.EnvCeleryTask() + ".NLP_edit_ASR",
		Lang:         "Go",
		RootId:       id,
		IgnoreResult: false,
		Retries:      0,
		Origin:       "Go-User",
	}

	properties := utils.CeleryProperties{
		Priority:      0,
		CorrelationId: id,
		ReplyTo:       id,
		DeliveryMode:  2,
		DeliveryInfo:  utils.CeleryPropertyDeliverInfo{RoutingKey: utils.CeleryRedisKey, Exchange: ""},
		BodyEncoding:  utils.BodyEncoding,
		DeliveryTag:   id,
	}

	task := utils.CeleryTask{Body: bodyBase64, ContentEncoding: utils.ContentEncoding, ContentType: utils.ContentType, Headers: headers, Properties: properties}

	jsonData, err := task.ToJson()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := r.client.LPush(ctx, utils.CeleryRedisKey, jsonData).Err(); err != nil {
		return err
	}
	return nil
}
func (r *Repository) MailMission(mail string) error {
	body := utils.CeleryBody{
		Position: []utils.CeleryBodyPosition{},
		Args: utils.CeleryMailBodyArgs{
			Mail: mail,
		},
	}
	bodyBase64, err := body.ToJsonBase64()
	if err != nil {
		return err
	}

	id := uuid.New().String()

	headers := utils.CeleryHeaders{
		Id:           id,
		Task:         utils.EnvCeleryTask() + ".Mail_sent",
		Lang:         "Go",
		RootId:       id,
		IgnoreResult: false,
		Retries:      0,
		Origin:       "Go-User",
	}

	properties := utils.CeleryProperties{
		Priority:      0,
		CorrelationId: id,
		ReplyTo:       id,
		DeliveryMode:  2,
		DeliveryInfo:  utils.CeleryPropertyDeliverInfo{RoutingKey: utils.CeleryRedisKey, Exchange: ""},
		BodyEncoding:  utils.BodyEncoding,
		DeliveryTag:   id,
	}

	task := utils.CeleryTask{Body: bodyBase64, ContentEncoding: utils.ContentEncoding, ContentType: utils.ContentType, Headers: headers, Properties: properties}

	jsonData, err := task.ToJson()
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := r.client.LPush(ctx, utils.CeleryRedisKey, jsonData).Err(); err != nil {
		return err
	}
	return nil
}
