package SE421.Assignment2;

import java.util.*;

public class song {
    private String name;
    private int length;
    private String album;
    private LinkedList<artist> artists = new LinkedList<artist>();
    private LinkedList<String> composers = new LinkedList<String>();

    public void print() {
        System.out.printf(name + " " + length + " " + album + " " + composers);
        for (artist artist : artists) {
            System.out.printf(artist.getArtistName());
        }
        System.out.println();
    }

    public LinkedList<String> getComposers() {
        return this.composers;
    }

    public void setComposers(String composer) {
        this.composers.add(composer);
    }

    public LinkedList<artist> getArtists() {
        return this.artists;
    }

    public void setArtists(artist artist) {
        this.artists.add(artist);
    }

    public int getLength() {
        return this.length;
    }

    public void setLength(int length) {
        this.length = length;
    }

    public String getName() {
        return this.name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getAlbum() {
        return this.album;
    }

    public void setAlbum(String album) {
        this.album = album;
    }

}
