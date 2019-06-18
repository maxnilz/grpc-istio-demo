const {EmojiRequest, EmojiResponse} = require('./emoji-without-annotations_pb.js');
const {EmojiServiceClient} = require('./emoji-without-annotations_grpc_web_pb.js');

var client = new EmojiServiceClient('http://' + window.location.hostname + ':31380/demo-server');
var editor = document.getElementById('editor');

window.insertEmojis = function() {
  var request = new EmojiRequest();
  request.setInputText(editor.innerText);
  client.insertEmojis(request, {}, (err, response) => {
    editor.innerText = response.getOutputText();
    window.focusEditor();
  });
};

window.focusEditor = function() {
  editor.focus();
  var range = document.createRange();
  range.selectNodeContents(editor);
  range.collapse(false);
  var sel = window.getSelection();
  sel.removeAllRanges();
  sel.addRange(range);
};

window.focusEditor();
