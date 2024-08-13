package database

import (
	"context"
	"fmt"
	"log"

	models "organum/Models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var clients *mongo.Client

func AddBookDB(eu *models.Book) (*mongo.InsertOneResult, error) {
	collection := client.Database("Organum").Collection("books")
	insertResult, err := collection.InsertOne(context.Background(), eu)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func ListBooksByStatusBD(t string) ([]models.Book, error) {
	collection := client.Database("Organum").Collection("books")
	filter := bson.M{"tag": t}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var books []models.Book
	for cursor.Next(context.Background()) {
		var book models.Book
		err := cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBookById(title string) (string, error) {
	collection := client.Database("Organum").Collection("books")
	filter := bson.M{"title": title}

	var book models.Book
	err := collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("Documento com título '%s' não encontrado", title)
		}
		return "", err
	}

	return book.ID, nil
}

func GetBookTitleByID(id string) (string, error) {
	collection := client.Database("Organum").Collection("books")
	filter := bson.M{"id": id}

	var book models.Book
	err := collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("Livro com ID '%s' não encontrado", id)
		}
		return "", err
	}

	return book.Title, nil
}

func GetAllPages(idbook string) (int, error) {
	collection := client.Database("Organum").Collection("books")
	filter := bson.M{"id": idbook}

	var allpage int
	var book models.Book

	err := collection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		fmt.Printf("Erro ao consultar:", err)
		return 0, err
	}

	allpage = book.Pages

	return allpage, nil

}

func UpdateBookTag(bookID, newTag string) (*mongo.UpdateResult, error) {
	collection := client.Database("Organum").Collection("books")
	filter := bson.M{"id": bookID}
	updateResult, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"tag": newTag}})
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func CreateReviewDB(review models.AddReview, newid string) (*mongo.InsertOneResult, error) {
	collection := client.Database("Organum").Collection("reviews")
	tt, _ := GetBookById(review.Title)
	rw := models.Review{
		ID:     newid,
		Opnion: review.Opnion,
		Score:  review.Score,
		IDbook: tt,
	}
	insertResult, err := collection.InsertOne(context.Background(), rw)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func GetReviewBD() ([]models.AddReview, error) {
	collection := client.Database("Organum").Collection("reviews")
	var reviews []models.AddReview

	all, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer all.Close(context.TODO())

	for all.Next(context.TODO()) {
		var rw models.Review
		if err := all.Decode(&rw); err != nil {
			return nil, err
		}

		title, err := GetBookTitleByID(rw.IDbook)
		if err != nil {
			return nil, err
		}
		get := models.AddReview{
			Title:  title,
			Opnion: rw.Opnion,
			Score:  rw.Score,
		}

		reviews = append(reviews, get)
	}

	if err := all.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}

func CreateQuoteBD(bookID, text string) (*mongo.UpdateResult, error) {
	collection := client.Database("Organum").Collection("quotes")
	filter := bson.M{"idbook": bookID}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		bookQuotes := models.Quote{
			IDBook: bookID,
			Quotes: []string{
				text,
			},
		}
		_, err := collection.InsertOne(context.Background(), bookQuotes)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Documento inserido com sucesso")
		return nil, nil
	}

	add := bson.M{"$push": bson.M{"quotes": text}}
	result, err := collection.UpdateOne(context.Background(), filter, add)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Citação adicionada com sucesso")
	return result, nil
}

func GetQuotesBD() ([]models.GetQuote, error) {
	collection := client.Database("Organum").Collection("quotes")
	var quotes []models.GetQuote

	all, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer all.Close(context.TODO())

	for all.Next(context.TODO()) {
		var quote models.Quote
		if err := all.Decode(&quote); err != nil {
			return nil, err
		}

		title, err := GetBookTitleByID(quote.IDBook)
		if err != nil {
			return nil, err
		}
		getQuote := models.GetQuote{
			Title:  title,
			Quotes: quote.Quotes,
		}

		quotes = append(quotes, getQuote)
	}

	if err := all.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}

func GetReadPages(idbook string) (int, error) {
	collection := client.Database("Organum").Collection("progress")
	filter := bson.M{"idbook": idbook}
	var progress models.Progress

	err := collection.FindOne(context.TODO(), filter).Decode(&progress)
	if err != nil {
		fmt.Printf("Erro ao consultar:", err)
		return 0, err
	}

	pages := progress.Pgl

	return pages, nil
}

func UpdateBookProgressDB(idbook string, pages int) (*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	collection := client.Database("Organum").Collection("progress")
	filter := bson.M{"idbook": idbook}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	all, err := GetAllPages(idbook)
	if err != nil {
		return nil, fmt.Errorf("Erro ao capturar todas as pags")
	}

	id := models.GenerateID()

	if count != 0 {
		var existing models.Progress
		err = collection.FindOne(context.Background(), filter).Decode(&existing)
		if err != nil {
			return nil, fmt.Errorf("Erro ao pesquisar progressos: %w", err)
		}

		newPages := existing.Pgl + pages

		update := bson.M{"$set": bson.M{"pgl": newPages}}
		result, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Printf("Erro ao atualizar progresso: %v", err)
			return nil, fmt.Errorf("Erro ao atualizar progresso: %w", err)
		}
	} else {
		progress := models.Progress{
			IDbook: idbook,
			ID:     id,
			Pgt:    all,
			Pgl:    pages,
		}
		_, err = collection.InsertOne(context.Background(), progress)
		if err != nil {
			log.Printf("Erro ao criar progresso: %v", err)
			return nil, fmt.Errorf("Erro ao criar progresso: %w", err)
		}
		fmt.Println("Documento inserido com sucesso")
		return nil, nil
	}

	return result, nil
}

func DeleteBookDB(idbook string) (*mongo.DeleteResult, error) {
	collection := client.Database("Organum").Collection("books")
	filter := bson.M{"id": idbook}
	DeleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return DeleteResult, nil
}

func DeleteProgressDB(idbook string) (*mongo.DeleteResult, error) {
	collection := client.Database("Organum").Collection("progress")
	filter := bson.M{"idbook": idbook}
	DeleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return DeleteResult, nil
}

func DeleteQuotesDB(idbook string) (*mongo.DeleteResult, error) {
	collection := client.Database("Organum").Collection("quotes")
	filter := bson.M{"idbook": idbook}
	DeleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return DeleteResult, nil
}

func DeleteReviewsDB(idbook string) (*mongo.DeleteResult, error) {
	collection := client.Database("Organum").Collection("reviews")
	filter := bson.M{"idbook": idbook}
	DeleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return DeleteResult, nil
}
