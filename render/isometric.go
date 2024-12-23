/*
Copyright (c) 2017 Lauris Bukšis-Haberkorns <lauris@nix.lv>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package render

import (
	"image"

	"github.com/disintegration/imaging"
	tiled "github.com/lafriks/go-tiled"
)

// IsometricRendererEngine represents isometric rendering engine.
type IsometricRendererEngine struct {
	m *tiled.Map
}

// Init initializes rendering engine with provided map options.
func (e *IsometricRendererEngine) Init(m *tiled.Map) {
	e.m = m
}

// GetFinalImageSize returns final image size based on map data.
func (e *IsometricRendererEngine) GetFinalImageSize() image.Rectangle {
	// In isometric rendering, the width and height will be different due to the staggered tile arrangement.
	// The final image width is affected by the number of tiles in x direction and the tile width.
	// The height will be affected by the number of tiles in y direction and the tile height.
	// width := (e.m.Width * e.m.TileWidth)
	// height := (e.m.Height * e.m.TileHeight)

	// return image.Rect(0, 0, width, height)

	return image.Rect(0, 0, e.m.Width*32, e.m.Height*32)

}

// RotateTileImage rotates provided tile layer.
func (e *IsometricRendererEngine) RotateTileImage(tile *tiled.LayerTile, img image.Image) image.Image {
	timg := img
	if tile.HorizontalFlip {
		timg = imaging.FlipH(timg)
	}
	if tile.VerticalFlip {
		timg = imaging.FlipV(timg)
	}
	if tile.DiagonalFlip {
		timg = imaging.FlipH(imaging.Rotate90(timg))
	}

	return timg
}

// GetTilePosition returns tile position in image.
func (e *IsometricRendererEngine) GetTilePosition(x, y, width, height int) image.Rectangle {
	// In isometric rendering, the formula is adjusted based on the tile's position in the isometric grid
	// screenX := (x - y) * e.m.TileWidth / 2
	// screenY := (x + y) * e.m.TileHeight / 2
	screenX := (x - y) * width / 2
	screenY := (x + y) * height / 4

	// Return a rectangle representing the tile's position and size in the final image
	// return image.Rect(screenX, screenY, screenX+e.m.TileWidth, screenY+e.m.TileHeight)
	// return image.Rect(screenX, screenY, screenX+32, screenY+32)

	return image.Rect(screenX, screenY, screenX+width, screenY+height)
}
