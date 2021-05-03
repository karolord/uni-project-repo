import java.awt.*;
import java.awt.geom.*;
import javax.swing.*;
import javax.swing.event.*;

public class ColorWheel extends JPanel {
    private final JSlider Slider;
    private Arc arc;`

    public ColorWheel() {
        setLayout(new BorderLayout());
        Slider = new JSlider(SwingConstants.HORIZONTAL, 0, 360, 0);
        Slider.setMajorTickSpacing(30);
        Slider.setPaintTicks(true);
        Slider.setPaintLabels(true);
        Slider.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                JSlider slider = (JSlider) e.getSource();
                arc.setAngle(slider.getValue());
            }
        });
        add(Slider, BorderLayout.SOUTH);
    }

    public class Arc extends JPanel {
        private double angle = 0;

        public void paintComponent(Graphics g) {
            super.paintComponent(g);
            Graphics2D arc = (Graphics2D) g.create();
            arc.setColor(Color.GREEN);
            arc.fill(new Arc2D.Double(10, 10, 300, 300, 90, 90, Arc2D.PIE));
            arc.setColor(Color.BLUE);
            arc.fill(new Arc2D.Double(10, 10, 300, 300, 180, 90, Arc2D.PIE));
            arc.setColor(Color.YELLOW);
            arc.fill(new Arc2D.Double(10, 10, 300, 300, 270, 90, Arc2D.PIE));
            arc.setColor(Color.RED);
            arc.fill(new Arc2D.Double(10, 10, 300, 300, 0, 90, Arc2D.PIE));
            arc.rotate(Math.toRadians(angle), 0, 0);

        }

        public void setAngle(double angle) {
            this.angle = angle;
            repaint();
        }

    }

    public static void main(String[] args) {
        JFrame frame = new JFrame();
        JPanel panel = new ColorWheel();
        panel.setBackground(Color.PINK);
        frame.add(panel);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setSize(330, 400);
        frame.setVisible(true);

    }
}