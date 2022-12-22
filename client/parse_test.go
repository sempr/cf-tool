package client

import (
	"fmt"
	"testing"
)

func TestSampleWithBr(t *testing.T) {
	var html = `
<div class="sample-tests">
	<div class="section-title">Examples</div>
	<div class="sample-test">
		<div class="input">
			<div class="title">Input</div>
			<pre>2
1 2 3 4<br /></pre>
		</div>
		<div class="output">
			<div class="title">Output</div>
			<pre>1<br /></pre>
		</div>
		<div class="input">
			<div class="title">Input</div>
			<pre>4<br />1 3 4 6 3 4 100 200<br /></pre>
		</div>
		<div class="output">
			<div class="title">Output</div>
			<pre>5<br /></pre>
		</div>
		<div class="input">
			<div class="title">Input</div>
			<pre>4<br /><br /><br /><br/><br /><br/>1 3 4 6 3 4 100 200<br /></pre>
		</div>
		<div class="output">
			<div class="title">Output</div>
			<pre>5</pre>
		</div>
		<div class="input">
			<div class="title">Input</div>
			<pre>0</pre>
		</div>
		<div class="output">
			<div class="title">Output</div>
			<pre>5</pre>
		</div>
	</div>
</div>`
	var data []byte = []byte(html)
	i, o, _ := findSample(data)

	for ii := range i {
		fmt.Println("---------->")
		fmt.Println(string(i[ii]))
		fmt.Println("<----")
		fmt.Println(string(o[ii]))
	}
	t.Error("")
}
