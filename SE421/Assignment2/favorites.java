import java.util.*;

public class favorites {
    private static LinkedList<artist> artists = new LinkedList<artist>();
    private static LinkedList<song> songs = new LinkedList<song>();

    public static LinkedList<artist> getArtists() {
        return artists;
    }

    public void setArtists(artist artist) {
        artists.add(artist);
    }

    public static LinkedList<song> getSongs() {
        return songs;
    }

    public void setSongs(song song) {
        songs.add(song);
    }

    public static void printFavorites() {
        for (int i = 1; i <= artists.size(); i++) {
            System.out.println("songs by " + artists.get(i - 1).getArtistName());
            System.out.println("Title\tLength\tAlbum\tComposers");
            artists.get(i - 1).getMediator().printSong();
            System.out.println();
        }
    }
}
