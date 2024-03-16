package window

import (
	"testing"

	"github.com/nomad-software/goat/example/image"
	"github.com/nomad-software/goat/image/store"
	"github.com/stretchr/testify/assert"
)

func TestWindow(t *testing.T) {
	win := New(nil)

	assert.Equal(t, "window", win.GetType())
	assert.Equal(t, "Toplevel", win.GetStyle())

	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}$`, win.GetID())
}

func TestWindowParent(t *testing.T) {
	win := New(nil)
	child := New(win)

	assert.Equal(t, "window", child.GetType())
	assert.Equal(t, "Toplevel", child.GetStyle())

	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}\.window-[A-Z0-9]{1,8}$`, child.GetID())
}

func TestWindowSize(t *testing.T) {
	win := New(nil)

	win.SetSize(250, 250)
	win.Update()

	assert.Equal(t, 250, win.GetWidth())
	assert.Equal(t, 250, win.GetHeight())
}

func TestWindowGeometry(t *testing.T) {
	win := New(nil)

	win.SetGeometry(350, 350, 150, 150)
	win.Update()

	assert.Equal(t, 350, win.GetWidth())
	assert.Equal(t, 350, win.GetHeight())

	assert.Equal(t, 150, win.GetXPos(false))
	assert.Equal(t, 187, win.GetYPos(false))
}

func TestWindowTitle(t *testing.T) {
	win := New(nil)
	win.SetTitle("foo")

	assert.Equal(t, "foo", win.GetTitle())
}

func TestWindowWaitForVisiblity(t *testing.T) {
	win := New(nil)

	win.SetSize(250, 250)
	win.WaitForVisibility()

	assert.Equal(t, 250, win.GetWidth())
	assert.Equal(t, 250, win.GetHeight())
}

func TestWindowFullScreen(t *testing.T) {
	win := New(nil)
	assert.False(t, win.IsFullScreen())

	win.SetFullScreen(true)
	win.WaitForVisibility()

	assert.True(t, win.IsFullScreen())
}

func TestWindowTopmost(t *testing.T) {
	win := New(nil)
	assert.False(t, win.IsTopmost())

	win.SetTopmost(true)
	win.WaitForVisibility()

	assert.True(t, win.IsTopmost())
}

func TestWindowIconfiy(t *testing.T) {
	win := New(nil)
	win.SetIconify(true)
	win.SetIconify(false)
}

func TestWindowMinMaxSize(t *testing.T) {
	win := New(nil)

	win.SetMinSize(100, 100)
	win.SetMaxSize(200, 200)

	win.SetSize(250, 250)
	win.Update()
	assert.Equal(t, 200, win.GetWidth())
	assert.Equal(t, 200, win.GetHeight())

	win.SetSize(50, 50)
	win.Update()
	assert.Equal(t, 100, win.GetWidth())
	assert.Equal(t, 100, win.GetHeight())
}

func TestWindowResizable(t *testing.T) {
	win := New(nil)

	res := win.GetResizeable()
	assert.True(t, res[0])
	assert.True(t, res[1])

	win.SetResizeable(false, false)

	res = win.GetResizeable()
	assert.False(t, res[0])
	assert.False(t, res[1])
}

func TestWindowIsAbove(t *testing.T) {
	win := New(nil)
	win.Update()

	child := New(win)
	child.Update()

	assert.True(t, child.IsAbove(win))
	assert.True(t, win.IsBelow(child))
}

func TestWindowIsBelow(t *testing.T) {
	win := New(nil)
	child := New(win)

	win.SetTopmost(true)
	win.Update()

	assert.True(t, child.IsBelow(win))
	assert.True(t, win.IsAbove(child))
}

func TestWindowIcon(t *testing.T) {
	store := store.New(image.FS)
	icons := store.GetImages("png/tkicon.png")

	win := New(nil)
	win.SetIcon(icons, false)
}
