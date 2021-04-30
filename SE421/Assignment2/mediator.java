import java.util.*;

public class mediator {
    public LinkedList<artist> artists = new LinkedList<artist>();
    public LinkedList<song> songs = new LinkedList<song>();

    public LinkedList<song> getSongs() {
        return this.songs;
    }

    public void setSongs(song song) {
        songs.add(song);
    }

    public LinkedList<artist> getArtists() {
        return this.artists;
    }

    public void setArtists(artist artist) {
        artists.add(artist);
    }

    public void printArtists() {
        for (int i = 1; i <= artists.size(); i++) {
            System.out.printf(artists.get(i - 1).getArtistName() + " ");
        }
    }

    public void printSong() {
        for (int i = 1; i <= songs.size(); i++) {
            System.out.printf(songs.get(i - 1).getName() + "\t" + songs.get(i - 1).getLength() + "\t"
                    + songs.get(i - 1).getAlbum() + "\t" + songs.get(i - 1).getComposers());
        }
    }

    public void deletesong(String s) {
        for (int i = 1; i <= songs.size(); i++) { 
            if (songs.get(i -  1 ).getName().equals(s)) {
                songs.remove(i - 1);
            }
        }
    }
}