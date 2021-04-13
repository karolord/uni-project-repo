package SE421.Assignment2;

import java.util.*;

public class mediator2 {
    public LinkedList<artist> artists = new LinkedList<artist>();
    public LinkedList<song> songs = new LinkedList<song>();

    public LinkedList<song> getSongs() {
        return this.songs;
    }

    public void setSongs(song song) {
        this.songs.add(song);
    }

    public LinkedList<artist> getArtists() {
        return this.artists;
    }

    public void setArtists(artist artist) {
        artists.add(artist);
    }
}
