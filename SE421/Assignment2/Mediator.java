package SE421.Assignment2;

import java.util.*;

public class mediator {
    public static LinkedList<artist> artists = new LinkedList<artist>();
    public static LinkedList<song> songs = new LinkedList<song>();
    public static LinkedList<composer> composers = new LinkedList<composer>();
    public static favorites favorites = new favorites();

    public static void getinputs() {
        Scanner input = new Scanner(System.in);
        for (int i = 0; i < 40; i++) {
            song song = new song();
            song artist = new artist();
            System.out.println("please enter the name of the song");
            song.setName(input.nextLine());
            System.out.println("please enter the song length in seconds");
            song.setLength(input.nextInt());
            System.out.println("please enter the album name");
            song.setAlbum(input.nextLine());
            mediator.songs.add(song);
            System.out.println("please enter the artist name");
            artist.set(input.nextLine());

        }
    }
}
