// This file contains types inside the test2 domain.

// Package app contains structs and implementations for this app.
package app

import (
	"github.com/Bitspark/go-bitnode/bitnode"
	"log"
)

// Entry description: A diary entry.
type Entry struct {
	// Date description:
	Date string `json:"date"`

	// EntryContent description:
	EntryContent string `json:"entryContent"`

	// Id description:
	Id string `json:"id"`

	// Title description:
	Title string `json:"title"`
}

// Tag description: Tag for an entry to categorize its content.
type Tag struct {
	// Name description:
	Name string `json:"name"`

	// Category description:
	Category string `json:"category"`

	// Id description:
	Id string `json:"id"`
}

// DOMAIN STRUCT

// Domain containing mainly wrappers for applications.
type Domain struct {
	Domain *bitnode.Domain
	Node   bitnode.Node
}

// NewDiary creates a new Diary instance.
func (test2 *Domain) NewDiary() (*Diary, error) {
	// Get the Diary sparkable from the domain.
	diarySpark, err := test2.Domain.GetSparkable("hub.testing.test2.Diary")
	if err != nil {
		log.Fatal(err)
	}

	// Remove docker implementation.
	delete(diarySpark.Implementation, "docker")

	// Prepare the Diary spark.
	diarySpk, err := test2.Node.PrepareSystem(bitnode.Credentials{}, *diarySpark)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the Diary.
	diary := &Diary{
		System: diarySpk,
	}

	return diary, nil
}
