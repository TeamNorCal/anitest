$(document).ready(function() {
  $("#startButton").on("click", anim.startAnimating);
  $("#stopButton").on("click", anim.stopAnimating);

  console.log("Calling init");
  $.ajax("/init")
  .done(function(data, status) {
    anim.startAnimating();
  })
  .fail(function(jqXHR, status, errorThrown) {
    console.error("Failed with status " + status + " and error " + errorThrown);
  });
});

var anim = {
  pending: false,
  missedFrameCount: 0,
  processFrame: function(frame) {
    frame.Universes.forEach(function(universe) {
      var data = universe.Data;
      var pixels = $("#universe" + universe.ID + " .pixel");
      data.forEach(function(pd, idx) {
        var color = "rgba(" + pd.R + ", " + pd.G + ", " + pd.B + ", " + pd.A + ")";
        pixels[idx].style["backgroundColor"] = color;
      });
    });
  },
  startAnimating: function() {
    if (!anim.timer) {
      console.log("Starting");
      anim.timer = setInterval(function() {
        //console.log("Getting frame");
        if (anim.pending) {
          console.log("Request still pending on frame tick; skipping");
          anim.missedFrameCount++;
          $("#missedFrameCount").text(anim.missedFrameCount);
        } else {
          anim.pending = true;
          $.getJSON("/getFrame", function(frame) {
            //console.log("Processing frame data", frame.Data[0]);
            anim.processFrame(frame);
            anim.pending = false;
          });
        }
      }, 33);
    }
  },
  stopAnimating: function() {
    if (anim.timer) {
      console.log("Stopping");
      clearInterval(anim.timer);
      delete anim.timer;
    }
  }
};
