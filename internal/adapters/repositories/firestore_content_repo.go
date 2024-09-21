package repositories

import (
	"ai-content-creator-agent/internal/domain/entities"
	"ai-content-creator-agent/internal/domain/interfaces"
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FirestoreContentRepository struct {
	Client *firestore.Client
	Ctx    context.Context
}

func NewContentRepository(ctx context.Context, projectID string) interfaces.ContentRepository {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}

	return &FirestoreContentRepository{
		Client: client,
		Ctx:    ctx,
	}
}

func (repo *FirestoreContentRepository) Save(content *entities.Content) error {
	_, err := repo.Client.Collection("contents").Doc(content.ID).Set(repo.Ctx, content)
	return err
}

func (repo *FirestoreContentRepository) Update(content *entities.Content) error {
	_, err := repo.Client.Collection("contents").Doc(content.ID).Set(repo.Ctx, content)
	return err
}

func (repo *FirestoreContentRepository) Get(companyID, contentID string) (*entities.Content, error) {
	doc, err := repo.Client.Collection("contents").Doc(contentID).Get(repo.Ctx)
	if err != nil {
		return nil, err
	}

	var content entities.Content
	err = doc.DataTo(&content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (repo *FirestoreContentRepository) GetAll(companyID string, limit int) ([]entities.Content, error) {
	var contents []entities.Content
	iter := repo.Client.Collection("contents").Where("company_id", "==", companyID).Limit(limit).Documents(repo.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var content entities.Content
		err = doc.DataTo(&content)
		if err != nil {
			return nil, err
		}

		contents = append(contents, content)
	}

	return contents, nil
}

func (repo *FirestoreContentRepository) Delete(contentID string) error {
	_, err := repo.Client.Collection("contents").Doc(contentID).Delete(repo.Ctx)
	return err
}

// GetPopular retrieves the most popular interms of likes content for the given company.
func (repo *FirestoreContentRepository) GetPopular(companyID string) (*entities.Content, error) {
	iter := repo.Client.Collection("contents").Where("company_id", "==", companyID).OrderBy("likes", firestore.Desc).Limit(1).Documents(repo.Ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, errors.New("no popular content found")
	}
	
	if err != nil {
		return nil, err
	}

	var content entities.Content
	err = doc.DataTo(&content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}