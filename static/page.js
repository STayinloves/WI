var editor = CodeMirror.fromTextArea(document.getElementById("code"), {
    theme: "solarized",
    keyMap: "vim",
    matchBrackets: true,
    indentUnit: 8,
    tabSize: 8,
    indentWithTabs: true,
    mode: "text/x-go"
})
function runCode() {
    app.waitServer = true, app.waitRun = false
    var data = {
        "version": 2,
        "body": editor.getValue()
    }
    console.log(data)
    axios.post('/compile', Qs.stringify(data))
        .then(function (res) {
                var error = res.data.Errors
                if (error !== "")
                    app.result = error
                else
                    app.result = res.data.Events[0].Message
                app.waitServer = false
            }
        ).catch(function (err) {
        console.log(err)
    })
}

var app = new Vue({
    el: '#pane',
    data: {
        waitServer: false,
        waitRun: true,
        result: "..."
    }
})
