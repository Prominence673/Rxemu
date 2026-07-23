package source

import (
	"context"
	"strings"
	"os/exec"
	"errors"
	"fmt"
	"encoding/json"
	"bytes"
	"time"
)

type Searcher interface{
	Search(
		ctx context.Context,
		query string,
		limit int,
	)([]Track, error)
 	Resolve(
        ctx context.Context,
        url string,
    ) (Track, error)
}

type YouTube struct{
	executable string
}

func NewYouTube() *YouTube{
	return &YouTube{
		executable: "yt-dlp",
	}
}

func (y *YouTube) Search(
	ctx context.Context,
	query string,
	limit int,
) ([]Track, error) {
	query = strings.TrimSpace(query)
	
	if query == "" {
		return nil, errors.New("search query cannot be empty")
	}

	if limit < 1 || limit > 25 {
		return nil, errors.New("search limit must be between 1 and 25")
	}
	
	searhCtx, cancel := context.WithTimeout(
		ctx,
		15*time.Second,
	)
	defer cancel()
	
	target := fmt.Sprintf(
		"ytsearch%d:%s",
		limit,
		query,
	)

	cmd := exec.CommandContext(
		searhCtx,
		y.executable,
		"--dump-single-json",
		"--flat-playlist",
		"--skip-download",
		"--no-warnings",
		target,
	)
	
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf(
			"yt-dlp search failed: %s: %w",
			strings.TrimSpace(stderr.String()),
			err,
		)
	}
	
	if errors.Is(searhCtx.Err(), context.DeadlineExceeded) {
		return nil, errors.New("YouTube search timed out")
	}
	
	var document searchDocument

	if err := json.Unmarshal(output, &document); err != nil {
		return nil, fmt.Errorf(
			"decode yt-dlp response: %w",
			err,
		)
	}
	
	tracks := make([]Track, 0, len(document.Entries))
	
	for _, entry := range document.Entries {
		artist := entry.Channel
		if artist == "" {
			artist = entry.Uploader
		}
	
		url := entry.WebpageURL
		if url == "" && entry.ID != "" {
			url = "https://www.youtube.com/watch?v=" + entry.ID
		}
	
		tracks = append(tracks, Track{
			ID:       entry.ID,
			Title:    entry.Title,
			Artist:   artist,
			Duration: entry.Duration,
			URL:      url,
		})
	}
	
	return tracks, nil
}

func (y *YouTube) Resolve(
    ctx context.Context,
    url string,
) (Track, error) {
    url = strings.TrimSpace(url)

    if url == "" {
        return Track{}, errors.New("URL cannot be empty")
    }

    resolveCtx, cancel := context.WithTimeout(
        ctx,
        15*time.Second,
    )
    defer cancel()

    cmd := exec.CommandContext(
        resolveCtx,
        y.executable,
        "--dump-single-json",
        "--skip-download",
        "--no-warnings",
        url,
    )

    var stderr bytes.Buffer
    cmd.Stderr = &stderr

    output, err := cmd.Output()
    if err != nil {
        if errors.Is(resolveCtx.Err(), context.DeadlineExceeded) {
            return Track{}, errors.New("resolve URL timed out")
        }

        if errors.Is(resolveCtx.Err(), context.Canceled) {
            return Track{}, errors.New("resolve URL canceled")
        }

        return Track{}, fmt.Errorf(
            "resolve URL with yt-dlp: %s: %w",
            strings.TrimSpace(stderr.String()),
            err,
        )
    }

    var entry searchEntry

    if err := json.Unmarshal(output, &entry); err != nil {
        return Track{}, fmt.Errorf(
            "decode yt-dlp response: %w",
            err,
        )
    }

    artist := entry.Channel
    if artist == "" {
        artist = entry.Uploader
    }

    trackURL := entry.WebpageURL
    if trackURL == "" {
        trackURL = url
    }

    return Track{
        ID:       entry.ID,
        Title:    entry.Title,
        Artist:   artist,
        Duration: entry.Duration,
        URL:      trackURL,
    }, nil
}