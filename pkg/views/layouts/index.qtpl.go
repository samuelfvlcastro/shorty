// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line ../../pkg/views/layouts/index.qtpl:1
package layouts

//line ../../pkg/views/layouts/index.qtpl:1
import "smashedbits.com/shorty/pkg/views/components/nav"

//line ../../pkg/views/layouts/index.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line ../../pkg/views/layouts/index.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line ../../pkg/views/layouts/index.qtpl:4
type Page interface {
//line ../../pkg/views/layouts/index.qtpl:4
	Title() string
//line ../../pkg/views/layouts/index.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line ../../pkg/views/layouts/index.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line ../../pkg/views/layouts/index.qtpl:4
	Body() string
//line ../../pkg/views/layouts/index.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line ../../pkg/views/layouts/index.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line ../../pkg/views/layouts/index.qtpl:4
	UserId() string
//line ../../pkg/views/layouts/index.qtpl:4
	StreamUserId(qw422016 *qt422016.Writer)
//line ../../pkg/views/layouts/index.qtpl:4
	WriteUserId(qq422016 qtio422016.Writer)
//line ../../pkg/views/layouts/index.qtpl:4
	UserEmail() string
//line ../../pkg/views/layouts/index.qtpl:4
	StreamUserEmail(qw422016 *qt422016.Writer)
//line ../../pkg/views/layouts/index.qtpl:4
	WriteUserEmail(qq422016 qtio422016.Writer)
//line ../../pkg/views/layouts/index.qtpl:4
}

//line ../../pkg/views/layouts/index.qtpl:12
func StreamBaseLayout(qw422016 *qt422016.Writer, p Page) {
//line ../../pkg/views/layouts/index.qtpl:12
	qw422016.N().S(`
<!DOCTYPE html>
<html data-theme="emerald">
<head>
  <title>`)
//line ../../pkg/views/layouts/index.qtpl:16
	p.StreamTitle(qw422016)
//line ../../pkg/views/layouts/index.qtpl:16
	qw422016.N().S(`</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <script src="dist/bundle.js"></script>
</head>
<body>
  `)
//line ../../pkg/views/layouts/index.qtpl:21
	nav.StreamRenderHeader(qw422016, p.UserId(), p.UserEmail())
//line ../../pkg/views/layouts/index.qtpl:21
	qw422016.N().S(`
  `)
//line ../../pkg/views/layouts/index.qtpl:22
	p.StreamBody(qw422016)
//line ../../pkg/views/layouts/index.qtpl:22
	qw422016.N().S(`
  `)
//line ../../pkg/views/layouts/index.qtpl:23
	nav.StreamRenderFooter(qw422016)
//line ../../pkg/views/layouts/index.qtpl:23
	qw422016.N().S(`
</body>
</html>
`)
//line ../../pkg/views/layouts/index.qtpl:26
}

//line ../../pkg/views/layouts/index.qtpl:26
func WriteBaseLayout(qq422016 qtio422016.Writer, p Page) {
//line ../../pkg/views/layouts/index.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../../pkg/views/layouts/index.qtpl:26
	StreamBaseLayout(qw422016, p)
//line ../../pkg/views/layouts/index.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line ../../pkg/views/layouts/index.qtpl:26
}

//line ../../pkg/views/layouts/index.qtpl:26
func BaseLayout(p Page) string {
//line ../../pkg/views/layouts/index.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line ../../pkg/views/layouts/index.qtpl:26
	WriteBaseLayout(qb422016, p)
//line ../../pkg/views/layouts/index.qtpl:26
	qs422016 := string(qb422016.B)
//line ../../pkg/views/layouts/index.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line ../../pkg/views/layouts/index.qtpl:26
	return qs422016
//line ../../pkg/views/layouts/index.qtpl:26
}