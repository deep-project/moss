package imagex

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/nfnt/resize"
)

// Image 图片处理工具
// 包括缩放、提取缩略图、计算宽高比
type Image struct {
	width  int
	height int

	// 图片缩放算法
	// NearestNeighbor: Nearest-neighbor插值
	// Bilinear：双线性插值
	// Bicubic：双三次插值
	// MitchellNetravali:Mitchell-Netravali插值
	// Lanczos2:Lanczos重采样，a=2
	// Lanczos3:Lanczos重采样，a=3
	interp resize.InterpolationFunction
}

func New() *Image {
	return &Image{}
}

func (i *Image) SetWidth(n int) *Image {
	i.width = n
	return i
}

func (i *Image) SetHeight(n int) *Image {
	i.height = n
	return i
}

func (i *Image) SetInterp(interp resize.InterpolationFunction) *Image {
	i.interp = interp
	return i
}

// Resize 缩放图片
// 宽或高有一项为0，则等比例缩放
func (i *Image) Resize(img image.Image) image.Image {
	return resize.Resize(uint(i.width), uint(i.height), img, i.interp)
}

// ResizeByte 缩放图片 by []byte
func (i *Image) ResizeByte(b []byte) ([]byte, error) {
	img, err := i.Byte2Image(b)
	if err != nil {
		return []byte{}, err
	}
	img = i.Resize(img)
	return i.Image2Byte(img), nil
}

// Thumbnail 生成缩略图
// 如果图片小于设置的尺寸，则返回原图片
func (i *Image) Thumbnail(img image.Image) image.Image {
	return resize.Thumbnail(uint(i.width), uint(i.height), img, i.interp)
}

func (i *Image) ThumbnailByte(b []byte) (_ []byte, err error) {
	img, err := i.Byte2Image(b)
	if err != nil {
		return
	}
	imgN := resize.Thumbnail(uint(i.width), uint(i.height), img, i.interp)
	return i.Image2Byte(imgN), nil
}

// Byte2Image []byte转image
func (i *Image) Byte2Image(b []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return img, nil
}

// Image2Byte image转[]byte
func (i *Image) Image2Byte(img image.Image) []byte {
	var buf bytes.Buffer
	jpeg.Encode(&buf, img, &jpeg.Options{Quality: 80})
	return buf.Bytes()
}

// CropByte 智能提取图片 by byte
// 按照尺寸，智能提取图片核心部分
func (i *Image) CropByte(b []byte) ([]byte, error) {
	img, err := i.Byte2Image(b)
	if err != nil {
		return []byte{}, err
	}
	if img, err = i.Crop(img); err != nil {
		return []byte{}, err
	}
	return i.Image2Byte(img), nil
}

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

// Crop 智能提取图片
// 按照尺寸，智能提取图片核心部分
func (i *Image) Crop(img image.Image) (image.Image, error) {
	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, err := analyzer.FindBestCrop(img, i.width, i.height)
	if err != nil {
		return nil, err
	}
	croppedImg := img.(subImager).SubImage(topCrop)
	if croppedImg.Bounds().Dx() > i.width { // 当前提取的图片大于设定的尺寸
		croppedImg = i.Resize(croppedImg) // 缩放
	}
	return croppedImg, nil
}

// ComputeScale 通过最大宽高 按比例计算新的宽高
// 如果未超过最大宽高，则原样返回
// 最大宽高设置0则不限制
func ComputeScale(width, height, maxWidth, maxHeight int) (int, int) {
	if maxWidth > 0 && width > maxWidth {
		var scale = float64(width) / float64(height)
		width = maxWidth
		height = int(float64(width) / scale)
	}
	if maxHeight > 0 && height > maxHeight {
		var scale = float64(height) / float64(width)
		height = maxHeight
		width = int(float64(height) / scale)
	}
	return width, height
}
