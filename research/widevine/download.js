'use strict';

const data = [];
const video = document.querySelector('video');

const recorder = new MediaRecorder(video.mozCaptureStream());

recorder.ondataavailable = function(event) {
   data.push(event.data);
};

recorder.onstop = function() {
   const blob = URL.createObjectURL(new Blob(data));
   console.log(blob);
};

video.onpause = function() {
   recorder.stop();
};

video.currentTime = 0;
recorder.start();
