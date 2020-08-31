import React from 'react';
import axios from 'axios';
import HomeScreen from './HomeScreen';

class App extends React.Component {
  constructor (props) {
    super(props)
    this.state = { 
      accessToken: null, 
      currentTrack: {
        album: {
          images: [{url: ""}]
        },
        name: "",
        artists: [{name: ""}],
        durationMs: 0,
        progressMs: 0,
        noData: false
      },
      lastTrack: {
        album: {
          images: [{url: ""}]
        },
        name: "",
        artists: [{name: ""}],
        durationMs: 0,
        progressMs: 0,
        noData: false
      }
    };
    this.updateTrackProgress = this.updateTrackProgress.bind(this);
    this.getCurrentlyPlayingTrack = this.getCurrentlyPlayingTrack.bind(this);
    this.getLastPlayedTrack = this.getLastPlayedTrack.bind(this);
  }

  componentDidMount() {
    const hash = window.location.hash
      .substring(1).split("&")
      .reduce(function(initial, item) {
        if (item) {
          var parts = item.split("=");
          initial[parts[0]] = decodeURIComponent(parts[1]);
        }
        return initial;
      }, {});
    window.location.hash = "";

    let _accessToken = hash.access_token;
    if (_accessToken) {
      this.setState({
        accessToken: _accessToken
      });
    }
  }

  updateTrackProgress() {
    if (this.state.accessToken && !this.state.currentTrack.noData) {
        this.getCurrentlyPlayingTrack(this.state.accessToken);
    }
  }
  
  getCurrentlyPlayingTrack(token) {

  }
  
  getLastPlayedTrack(token) {
  
  }

  render() {
    return (
        <HomeScreen />
    );
  }
}

export default App;
