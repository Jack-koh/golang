$(function() {
    if(!window.EventSource) {
        alert("No EventSource!")
        return;
    }

    let $chatlog = $("#chat-log")
    let $chatmsg = $("#chat-msg")


    let isBlank = (string) => string == null || string.trim() === ""
    let username;

    while(isBlank(username)) {
        username = prompt("What's your name")
        if(!isBlank(username)) {
            $("#user-name").html(`<b>${username}<b/>`)
        }
    }

    $("#input-form").on("submit", function(e) {
        e.preventDefault()
        $.post("/messages", {
            msg: $chatmsg.val(),
            name: username
        })
        $chatmsg.val("")
        $chatmsg.focus()
    })

    const addMessage = (data) => {
        let text = "";
        if(!isBlank(data.name)) {
           text = `<strong>${data.name}:</strong> `;
        }
        text += data.msg
        $chatlog.prepend(`<div><span>${text}</span></div>`)
    }

    const es = new EventSource("/stream")
    es.onopen = function(e) {
        $.post('users/', { name: username })
    }
    es.onmessage = function(e) {
        const msg = JSON.parse(e.data)
        addMessage(msg)
    }
    window.onbeforeunload = function() {
        $.ajax({
            url: "/users?username=" + username,
            type: "DELETE"
        })
        es.close()
    }
})