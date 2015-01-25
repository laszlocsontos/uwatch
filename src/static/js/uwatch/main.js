(function ($) {
  var FIELD_YOUTUBE_CONVERT = '#youTube-convert';
  var FIELD_YOUTUBE_PLAYER = 'youTube-player';
  var FIELD_YOUTUBE_TITLE = '#youTube-title';

  var FIELD_YOUTUBE_URL = '#youTube-url';
  var FIELD_YOUTUBE_URL_CHECK_INTERVAL = 250;
  var FIELD_YOUTUBE_URL_MSG = '#youTube-url-msg';

  var YOUTUBE_API = '<script src="https://www.youtube.com/iframe_api" />';
  var YOUTUBE_LONG_VIDEO_URL = "api/long_video_url"

  var YOUTUBE_URL_FIRST = 'http://www.youtube.com/watch?v=pNuZIZOya78'
  var YOUTUBE_URL_REGEX = /http.+youtube\.com\/watch\?v\=(\w+)/;

  var UWatch = function() {
    var _this = this;

    _this.youTubeConvert = $(FIELD_YOUTUBE_CONVERT);
    _this.youTubeConvertForm = $('youTube-convert-fm');

    _this.youTubeConvertButton = $('youTube-convert-btn');

    _this.youTubeConvertButton.on('click', function(event) {
      _this.youTubeConvertForm.submit();
    });

    _this.youTubeTitle = $(FIELD_YOUTUBE_TITLE);

    _this.youTubeUrl = $(FIELD_YOUTUBE_URL);
    _this.youTubeUrl.val(YOUTUBE_URL_FIRST);

    _this.youTubeUrlMsg = $(FIELD_YOUTUBE_URL_MSG);
    _this.youTubeUrlMsg.hide();

    _this.createYouTubeAPI(null);
  };

  UWatch.prototype = {
    constructor: UWatch,

    checkUrl: function(url) {
      var youTubeId = this.getYouTubeId(url);

      if (youTubeId) {
        this.getMetaData(youTubeId);
      }
      else {
        if (console && console.log) {
           console.log(this);
        }

        this.disableControls();
        this.youTubeUrlMsg.show();
      }
    },

    createYouTubeAPI: function() {
      var _this = this;

      var youTubeId = _this.getYouTubeId(YOUTUBE_URL_FIRST);

      var youTubeEvents = {
        onReady: $.proxy(_this.onYouTubePlayerReady, _this),
        onStateChange: $.proxy(_this.onYouTubePlayerStateChange, _this)
      };

      window.onYouTubeIframeAPIReady = function() {
        _this.youTubePlayer = new YT.Player(FIELD_YOUTUBE_PLAYER, {
          /*height: '390', width: '640',*/
          videoId: youTubeId,
          events: youTubeEvents
        });

        if (console && console.log) {
          console.log('YouTube API has just got initialized.');
        }
      };

      var firstScriptTag = $('script').first();

      firstScriptTag.before(YOUTUBE_API);
    },

    detectChange: function(input, handler) {
      var old = input.attr('data-old-value');
      var current = input.val();

      if (old !== current) { 
        if (typeof old != 'undefined') { 
          handler.call(this, current);
        }

        input.attr('data-old-value', current);
      }
    },

    disableControls: function() {
      this.youTubeConvert.children().prop('disabled', true);
      $('#' + FIELD_YOUTUBE_PLAYER).hide();
    },

    enableControls: function() {
      this.youTubeConvert.children().prop('disabled', false);
      $('#' + FIELD_YOUTUBE_PLAYER).show();
    },

    getMetaData: function(youTubeId) {
      var _this = this;

      // TODO generalize URL parsing
      var url = YOUTUBE_LONG_VIDEO_URL + '/YouTube/' + youTubeId;

      $.get(url)
        .done(function(result) {
          _this.setMetaData(result);

          _this.enableControls();
          _this.youTubeUrlMsg.hide();
        })
        .fail(function(result) {
          if (console && console.log) {
            console.log(result);
          }

          _this.disableControls();

          // TODO show actual error message
          _this.youTubeUrlMsg.show();
        });
    },

    getYouTubeId: function(url) {
      var match = url.match(YOUTUBE_URL_REGEX);

      if (match && match[1]) {
        return match[1];
      }

      return null;
    },

    onConvert: function() {
      
    },

    onSliderChange: function(value) {
      if (console && console.log) {
        console.log(value);
      }

      this.youTubePlayer.seekTo(value, true);
      this.youTubePlayer.playVideo();
    },

    onYouTubePlayerReady: function() {
      var _this = this;

      if (console && console.log) {
        console.log('onYouTubePlayerReady');
      }

      _this.checkUrl(_this.youTubeUrl.val());

      setInterval(function() {
        _this.detectChange(_this.youTubeUrl, _this.checkUrl);
      }, FIELD_YOUTUBE_URL_CHECK_INTERVAL);
    },

    onYouTubePlayerStateChange: function() {
      if (console && console.log) {
        console.log('onYouTubePlayerStateChange');
      }
    },

    setMetaData: function(metaData) {
      var metaData = JSON.parse(metaData)

      var videoId = (metaData && metaData["VideoId"]) ? metaData["VideoId"] : ''
      var title = (metaData && metaData.Title) ? metaData.Title : '&nbsp;'

      if (title.length > 25) {
        title = title.substring(0, 25) + " ... "
      }

      // if (console && console.log) {
      //   console.log(metaData)
      //   console.log(videoId)
      //   console.log(title)
      // }

      this.youTubeTitle.text(title)

      this.youTubePlayer.loadVideoById({
        videoId: videoId,
        startSeconds: 0,
        suggestedQuality: 'small'
      });
    },

  };

  new UWatch();
})(jQuery);