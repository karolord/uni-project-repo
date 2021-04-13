package SE421.Assignment2;

import java.util.*;

public class Main {

    public static void main(String[] args) {
        LinkedList<song> songs = new LinkedList<song>();
        getinputs(songs);

    }

    public static void getinputs(LinkedList<song> songs) {
        Scanner input = new Scanner(System.in);
        for (int i = 0; i < 40; i++) {
            song song = new song();
            System.out.println("please enter the name of the song");
            song.setName(input.nextLine());
            /*
             * System.out.println("please enter the song length in seconds");
             * song.setLength(input.nextInt());
             * System.out.println("please enter the album name");
             * song.setAlbum(input.nextLine());
             */
            songs.add(song);
            System.out.println("Enter 0 to stop adding songs and any other nummber to continue");
            int x = input.nextInt();
            if (x == 0)
                break;

        }

    }
}
