public class addartist implements chain {
    private chain nextinChain;
    Boolean bool = true;

    public void setnextchain(chain nextchain) {
        this.nextinChain = nextchain;
    }

    public void addfavorite(String favorite) {
        for (int i = 1; i <= app.allartists.size(); i++) {
            if (app.allartists.get(i - 1).getArtistName().equals(favorite)) {
                app.favorite.setArtists(app.allartists.get(i - 1));
                bool = false;
            }
        }
        if (bool) {
            System.out.println("Error invalid input");
        }
    }
}
