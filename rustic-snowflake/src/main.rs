use raylib::prelude::*;

fn main() {
    let (screen_width, screen_height) = (1920, 1080);

    let (mut rl, thread) = raylib::init()
        .size(screen_width, screen_height)
        .title("Hello, World")
        .build();

    let centre = Vector2 {
        x: screen_width as f32 / 2.0,
        y: screen_height as f32 / 2.0
    };

    // println!("{}", PI);

    while !rl.window_should_close() {
        let mut d = rl.begin_drawing(&thread);

        d.draw_fps(30, 20);
        d.clear_background(Color::BLACK);
        draw_snowflakes(&mut d, centre, 4, 7, 200.0, 10.0);
    }
}

// struct Line {
//     start: Vector2,
//     end: Vector2,
//     colour: Color,
// }

fn draw_snowflakes(d: &mut RaylibDrawHandle, c: Vector2, levels: i8, branches: i8, length: f32, thickness: f32) {
    let branch_angle = 2.0 * PI / (branches as f64);
    let colour: Color;
    match levels{
        4=>colour = Color::RAYWHITE,
        3=>colour = Color::RED,
        2=>colour = Color::YELLOW,
        1=>colour = Color::ORANGE,
        _=>colour = Color::WHITE,
    }

    if levels > 0 {
        for branch in 0..branches {
            // println!("{}", branch);
            let line = Vector2 {
                x: c.x + (branch_angle as f32 * branch as f32).cos() * length,
                y: c.y + (branch_angle as f32 * branch as f32).sin() * length,
            };
            d.draw_line_ex(c, line, thickness, colour);
            // draw_snowflakes(d, line, levels - 1, branches, thickness/2.0);
            draw_snowflakes(d, line, levels - 1, branches, length/1.5, thickness/2.0);
        }
    }
}
