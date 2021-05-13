import java.awt.Graphics;

import javax.swing.JFrame;
import javax.swing.JPanel;
public class GuiAir extends JPanel{
    @Override
    public void paintComponent(Graphics g){
        super.paintComponent(g);
        g.drawArc(0, 150, 600, 225, 90, 180);
        g.drawLine(275, 150, 1000, 150);
        g.drawLine(275, 375, 1000, 375);
        g.drawLine(400, 150, 550, 50);
        g.drawLine(650, 150, 700, 50);
        g.drawLine(400, 375, 550, 475);
        g.drawLine(650, 375, 700, 475);
    }
    public static void main(String[] args) {
        JFrame frame = new JFrame();
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        GuiAir arc = new GuiAir();
        frame.add(arc);
        frame.setSize(1000,600);
        frame.setVisible(true);
    }
    
}

