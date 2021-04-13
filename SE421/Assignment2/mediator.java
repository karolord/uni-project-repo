package SE421.Assignment2;

import java.util.*;
import java.util.Iterator;

public class mediator {
    public static LinkedList<artist> artists = new LinkedList<artist>();
    public static LinkedList<song> songs = new LinkedList<song>();
    public static favorites favorite = new favorites();

    public static void getInputs() {
        Scanner input = new Scanner(System.in);
        while (true) {
            song song = new song();
            Iterator<artist> it = artists.iterator();
            System.out.println("please enter the name of the song");
            song.setName(input.nextLine());
            System.out.println("please enter the song length in seconds");
            song.setLength(input.nextInt());
            System.out.println("please enter the album name");
            song.setAlbum(input.nextLine());
            songs.add(song);
            while (true) {
                System.out.println("please enter the composer name");
                String composer = input.nextLine();
                System.out.println("If there is another composer press 0 else input another number");
                int x = input.nextInt();
                if (x != 0)
                    break;
            }
            while (true) {
                artist artist = new artist();
                System.out.println("please enter the artist name");
                artist.setArtistName(input.nextLine());
                song.setArtists(artist);
                for (int j = 1; j <= artists.size(); j++) {
                    if (!(artists.get(j - 1).getArtistName().equals(artist.getArtistName()))) {
                        artist.setSongs(song);
                        artists.add(artist);
                    } else {
                        artists.get(j - 1).setSongs(song);
                    }
                }
                System.out.println("If there is another artist press 0 else input another number");
                int x = input.nextInt();
                if (x != 0)
                    break;
            }
            System.out.println("If there is another song press 0 else input another number");
            int x = input.nextInt();
            if (x != 0)
                break;
        }
    }
}
