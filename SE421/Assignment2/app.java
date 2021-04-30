/*
    Karo Rasool kk19046@auis.edu.krd
    Kosar Aziz kn18011@auis.edu.krd
*/

import java.util.*;

public class app {
    public static LinkedList<artist> allartists = new LinkedList<artist>();
    public static LinkedList<song> allsongs = new LinkedList<song>();
    public static favorites favorite = new favorites();
    static Scanner input = new Scanner(System.in);

    public static void main(String[] args) {
        getInputs();
        menu();
    }

    public static void menu() {
        while (true) {
            System.out.println("1.Print table of all songs");
            System.out.println("2.Print table of a specific artist");
            System.out.println("3.Delete a song");
            System.out.println("4.add favorites");
            System.out.println("5.print your favorite arists songs");
            System.out.println("6.Quit");
            System.out.println(": ");
            int x = input.nextInt();
            input.nextLine();
            switch (x) {
            case 1:
                printall();
                break;
            case 2:
                printArtist();
                break;
            case 3:
                delete();
                break;
            case 4:
                favorites();
                break;
            case 5:
                favorites.printFavorites();
                break;
            default:
                return;
            }
        }
    }

    public static void printall() {
        System.out.println("Title\tLength\tAlbum\tComposers\tArtists");
        for (int i = 1; i <= allsongs.size(); i++) {
            allsongs.get(i - 1).printSong();
        }
    }

    public static void printArtist() {
        System.out.println("Please enter the artist name:");
        String name = input.nextLine();
        for (int i = 1; i <= allartists.size(); i++) {
            if (allartists.get(i - 1).getArtistName().equals(name)) {
                System.out.println("the songs by " + allartists.get(i - 1).getArtistName());
                System.out.println("Title\tLength\tAlbum\tComposers");
                allartists.get(i - 1).getMediator().printSong();
                System.out.println();
            }
        }
    }

    public static void delete() {
        System.out.println("Please enter the song name:");
        String name = input.nextLine();
        LinkedList<artist> tmpartists = new LinkedList<artist>();
        tmpartists = (LinkedList<artist>) allartists.clone();
        LinkedList<song> tmpsongs = new LinkedList<song>();
        tmpsongs = (LinkedList<song>) allsongs.clone();
        for (int x = 1; x <= allsongs.size(); x++) {
            if (allsongs.get(x - 1).getName().equals(name)) {
                for (int i = 1; i <= allsongs.get(x - 1).getArtist().getArtists().size(); i++) {
                    for (int j = 1; j <= allartists.size(); j++) {
                        if (allartists.get(j - 1).getArtistName()
                                .equals(allsongs.get(x - 1).getArtist().getArtists().get(i - 1).getArtistName()))
                            allartists.get(j - 1).getMediator().deletesong(name);
                    }
                }
                allsongs.remove(x - 1);
            }
        }
        System.out.println("If you want to undo enter 0 else enter another number:");
        if (input.nextInt() == 0) {
            allartists = (LinkedList<artist>) tmpartists.clone();
            allsongs = (LinkedList<song>) tmpsongs.clone();
        }
    }

    public static void getInputs() {
        while (true) {
            boolean bool = true;
            song song = new song();
            System.out.println("please enter the name of the song:");
            String name = input.nextLine();
            song.setName(name);
            System.out.println("please enter the song length in seconds:");
            int length = input.nextInt();
            input.nextLine();
            song.setLength(length);
            System.out.println("please enter the album name:");
            String AlbumName = input.nextLine();
            song.setAlbum(AlbumName);
            while (true) {
                System.out.println("please enter the composer name:");
                String composer = input.nextLine();
                song.setComposers(composer);
                System.out.println("If there is another composer press 0 else input another number:");
                int x = input.nextInt();
                input.nextLine();
                if (x != 0)
                    break;
            }
            while (true) {
                artist artist = new artist();
                System.out.println("please enter the artist name:");
                artist.setArtistName(input.nextLine());
                song.setArtist(artist);
                for (int j = 1; j <= allartists.size(); j++) {
                    if (allartists.get(j - 1).getArtistName().equals(artist.getArtistName())) {
                        allartists.get(j - 1).setSongs(song);
                        bool = false;
                    }
                }
                if (bool) {
                    artist.setSongs(song);
                    allartists.add(artist);
                }
                System.out.println("If there is another artist press 0 else input another number:");
                int x = input.nextInt();
                input.nextLine();
                if (x != 0)
                    break;
            }
            allsongs.add(song);
            System.out.println("If there is another song press 0 else input another number:");
            int x = input.nextInt();
            input.nextLine();
            if (x != 0)
                break;
        }
    }

    public static void favorites() {
        chain c1 = new addsong();
        chain c2 = new addartist();
        c1.setnextchain(c2);
        System.out.println("Please enter the name you want to add to favorites:");
        c1.addfavorite(input.nextLine());
    }

}
