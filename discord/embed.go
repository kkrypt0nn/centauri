package discord

import "time"

// Embed represents embedded content in a message (discord.Message)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-structure
type Embed struct {
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Timestamp   *time.Time     `json:"timestamp,omitempty"`
	Color       int            `json:"color,omitempty"`
	Footer      *EmbedFooter   `json:"footer,omitempty"`
	Image       *EmbedResource `json:"image,omitempty"`
	Thumbnail   *EmbedResource `json:"thumbnail,omitempty"`
	Video       *EmbedResource `json:"video,omitempty"`
	Provider    *EmbedProvider `json:"provider,omitempty"`
	Author      *EmbedAuthor   `json:"author,omitempty"`
	Fields      []EmbedField   `json:"fields,omitempty"`
}

// NewEmbedBuilder creates a new embed and facilitates the creation of embed structures (discord.Embed)
func NewEmbedBuilder() Embed {
	return Embed{}
}

// SetTitle sets the title of the embed
func (e Embed) SetTitle(title string) Embed {
	e.Title = title
	return e
}

// SetDescription sets the description of the embed
func (e Embed) SetDescription(description string) Embed {
	e.Description = description
	return e
}

// SetURL sets the URL of the embed
func (e Embed) SetURL(url string) Embed {
	e.URL = url
	return e
}

// SetTimestamp sets the URL of the embed
func (e Embed) SetTimestamp(time time.Time) Embed {
	e.Timestamp = &time
	return e
}

// SetColor sets the color of the embed
func (e Embed) SetColor(color int) Embed {
	e.Color = color
	return e
}

// SetFooter sets the footer of the embed
func (e Embed) SetFooter(text, iconURL string) Embed {
	e.Footer = &EmbedFooter{
		Text:    text,
		IconURL: iconURL,
	}
	return e
}

// SetImage sets the image of the embed
func (e Embed) SetImage(url string) Embed {
	e.Image = &EmbedResource{
		URL: url,
	}
	return e
}

// SetThumbnail sets the thumbnail of the embed
func (e Embed) SetThumbnail(url string) Embed {
	e.Thumbnail = &EmbedResource{
		URL: url,
	}
	return e
}

// SetAuthor sets the author of the embed
func (e Embed) SetAuthor(name, url, iconURL string) Embed {
	e.Author = &EmbedAuthor{
		Name:    name,
		URL:     url,
		IconURL: iconURL,
	}
	return e
}

// SetFields sets the fields of the embed
func (e Embed) SetFields(fields []EmbedField) Embed {
	e.Fields = fields
	return e
}

// AddField adds a field to the embed
func (e Embed) AddField(name, value string, inline bool) Embed {
	e.Fields = append(e.Fields, EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
	return e
}

// EmbedFooter represents footer information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedResource represents resource information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-thumbnail-structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedResource struct {
	URL          string `json:"url"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
}

// EmbedProvider represents the provider information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure
type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

// EmbedAuthor represents the author information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure
type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedField represents the field information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
