package SE421.Assignment2;

import java.util.*;

public class Main {
    public static void main(String[] args) {

        song songs[] = new song[40];
        getinputs(songs);
        for (int i = 0; i < 5; i++) {
            System.out.println(songs[i].getName());
        }
    }

    public static void getinputs(song songs[]) {
        Scanner input = new Scanner(System.in);
        for (int i = 0; i < 5; i++) {
            songs[i] = new song();
            System.out.println("please enter the name of the song");
            songs[i].setName(input.nextLine());

            System.out.println("please enter the song length in seconds");
            songs[i].setLength(input.nextInt());
            System.out.println("please enter the album name");
            songs[i].setAlbum(input.nextLine());

        }

    }
}
