package supabase

import (
	"errors"
	"os"

	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func Init() error {
	apiURL := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_KEY")

	if apiURL == "" || apiKey == "" {
		return errors.New("SUPABASE_URL or SUPABASE_KEY is missing in .env")
	}

	var err error
	Client, err = supabase.NewClient(apiURL, apiKey, nil)
	if err != nil {
		return err
	}

	return nil
}