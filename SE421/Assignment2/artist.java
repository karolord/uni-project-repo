package SE421.Assignment2;

import java.util.*;

public class artist {
    private String artistName;
    private LinkedList<song> songs = new LinkedList<song>();

    public LinkedList<song> getSongs() {
        return this.songs;
    }

    public void setSongs(song song) {
        this.songs.add(song);
    }

    public String getArtistName() {
        return this.artistName;
    }

    public void setArtistName(String artistName) {
        this.artistName = artistName;
    }

}
