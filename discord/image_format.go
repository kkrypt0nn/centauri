package discord

// ImageFormat represents the different image formats Discord accepts
// https://discord.com/developers/docs/reference#image-formatting-image-formats
type ImageFormat string

const (
	ImageFormatJPG    ImageFormat = "jpg"
	ImageFormatJPEG   ImageFormat = "jpeg"
	ImageFormatPNG    ImageFormat = "png"
	ImageFormatWebP   ImageFormat = "webp"
	ImageFormatGIF    ImageFormat = "gif"
	ImageFormatLottie ImageFormat = "json"
)
