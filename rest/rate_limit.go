package rest

import (
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// See https://discord.com/developers/docs/topics/rate-limits

// Bucket represents a rate limit bucket
type Bucket struct {
	sync.Mutex
	Key       string
	Remaining int
	Limit     int
	ResetAt   time.Time
}

// NewBucket returns a new rest.Bucket based on a key
func NewBucket(key string) *Bucket {
	return &Bucket{
		Key:       key,
		Remaining: 1,
	}
}

// RateLimiter represents the rate limiter holding a map of rest.Bucket and responsible to sleep when someone is doing too many requests
type RateLimiter struct {
	sync.Mutex
	Buckets       map[string]*Bucket
	GlobalResetAt time.Time
	MaxRetries    int
}

// NewRateLimiter returns a new rest.RateLimiter
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		Buckets:       make(map[string]*Bucket),
		GlobalResetAt: time.Unix(0, 0),
		MaxRetries:    5,
	}
}

// GetBucket returns a rest.Bucket based on the given key (URL without query parameters)
func (r *RateLimiter) GetBucket(key string) *Bucket {
	r.Lock()
	defer r.Unlock()

	// First check if a bucket for the given key is already existing
	if bucket, ok := r.Buckets[key]; ok {
		return bucket
	}

	// Otherwise create a bucket and add it to the map of buckets
	bucket := NewBucket(key)
	r.Buckets[key] = bucket
	return bucket
}

// LockBucket locks a rest.Bucket and waits if needed until it is no longer rate limited
func (r *RateLimiter) LockBucket(key string) *Bucket {
	// Get and lock the bucket
	bucket := r.GetBucket(key)
	bucket.Lock()

	// Wait the time required if it has no remaining requests and check for global rate limit as well
	if now := time.Now(); bucket.Remaining < 1 && bucket.ResetAt.After(now) {
		time.Sleep(bucket.ResetAt.Sub(now))
	} else if now := time.Now(); r.GlobalResetAt.After(now) {
		time.Sleep(r.GlobalResetAt.Sub(now))
	}

	// Consume the bucket
	if bucket.Remaining > 0 {
		bucket.Remaining--
	}
	return bucket
}

// UnlockBucket unlocks a rest.Bucket and sets the new values about the rate limits
func (r *RateLimiter) UnlockBucket(bucket *Bucket, headers http.Header) error {
	defer bucket.Unlock()

	// If we don't give headers, we don't need to calculate the new rate limits
	if headers == nil {
		return nil
	}

	// Set how many remaining requests we can make
	remaining := headers.Get("X-RateLimit-Remaining")
	if remaining != "" {
		parsed, err := strconv.ParseInt(remaining, 10, 64)
		if err != nil {
			return err
		}
		bucket.Remaining = int(parsed)
	}

	// Set the new reset and global reset times
	reset := headers.Get("X-RateLimit-Reset")
	resetAfter := headers.Get("X-RateLimit-Reset-After")
	global := headers.Get("X-RateLimit-Global")

	if resetAfter != "" {
		parsed, err := strconv.ParseFloat(resetAfter, 64)
		if err != nil {
			return err
		}
		integer, frac := math.Modf(parsed)

		resetAt := time.Now().Add(time.Duration(integer)*time.Second + time.Duration(frac*1000)*time.Millisecond)
		if global != "" {
			r.GlobalResetAt = resetAt
		} else {
			bucket.ResetAt = resetAt
		}
	} else if reset != "" {
		parsed, err := strconv.ParseFloat(reset, 64)
		if err != nil {
			return err
		}
		integer, frac := math.Modf(parsed)

		// Add some extra time to prevent encountering a rate limit again
		resetAt := time.Unix(int64(integer), int64(frac*1000*1000*1000)).Add(250 * time.Millisecond)
		bucket.ResetAt = resetAt
	}

	return nil
}
