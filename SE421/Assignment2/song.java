package SE421.Assignment2;

import java.util.*;

public class song {
    private String name;
    private int length;
    private String album;
    private mediator2 artists;
    private LinkedList<String> composers = new LinkedList<String>();

    public void print() {
        System.out.printf(name + " " + length + " " + album + " " + composers);
        System.out.println();
    }

    public void addArtists(artist artist) {
        mediator2 artists = new mediator2();
        mediator2.setArtists(artist);
    }

    public LinkedList<String> getComposers() {
        return this.composers;
    }

    public void setComposers(String composer) {
        this.composers.add(composer);
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
