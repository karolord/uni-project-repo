public class addsong implements chain {
    private chain nextinChain;
    Boolean bool = true;

    public void setnextchain(chain nextchain) {
        this.nextinChain = nextchain;
    }

    public void addfavorite(String favorite) {
        for (int i = 1; i <= app.allsongs.size(); i++) {
            if (app.allsongs.get(i - 1).getName().equals(favorite)) {
                app.favorite.setSongs(app.allsongs.get(i - 1));
                bool = false;
            }
        }
        if (bool) {
            nextinChain.addfavorite(favorite);
        }
    }

}
