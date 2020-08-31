import React from 'react';
import { Component } from "react";
import { Typography, Button } from '@material-ui/core';
import styled from 'styled-components';
import { authEndpoint, clientId, redirectUri, scopes } from "./config";

const HomeScreen = props => {
    return (
        <CenteredDiv>
        <Typography variant="h1">Notera</Typography>
        <Typography variant="h3">A Spotify Annotator</Typography>
        <p></p>
        <Typography variant="body1">Bookmark your Spotify tracks at specific timestamps and annotate them with your notes, or browse through other people's public annotations</Typography>
        <p></p>
        <Button variant="contained" color="secondary" 
            href={`${authEndpoint}?client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scopes.join(
                "%20"
            )}&response_type=token&show_dialog=true`}
        >Login to Spotify</Button>
        </CenteredDiv>
    );
  }


const CenteredDiv = styled.div`
flex: 1;
position: relative;
display: flex;
flex-direction: column;
justify-content: center;
align-items: center;
`;

  export default HomeScreen;