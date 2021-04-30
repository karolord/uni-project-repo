import java.util.*;

public class song {
    private String name;
    private int length;
    private String album;
    private LinkedList<String> composers = new LinkedList<String>();
    private mediator mediator = new mediator();

    public void printSong() {
        System.out.printf(name + "\t" + length + "\t" + album + "\t" + composers + "\t\t");
        mediator.printArtists();
        System.out.println();
    }

    public mediator getArtist() {
        return this.mediator;
    }

    public void setArtist(artist artists) {
        mediator.setArtists(artists);
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
