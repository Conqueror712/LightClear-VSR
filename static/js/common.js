function upload() {
    var uploadInput = document.getElementById("upload-input");
    var file = uploadInput.files[0];
    // Code to send file to server using Go backend
  }
  
  function download() {
    var downloadLink = document.getElementById("download-link");
    // Code to get file from server using Go backend
    downloadLink.href = 'file url from server';
  }
  