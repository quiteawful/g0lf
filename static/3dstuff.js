var levelstr = '{"Dim":{"X": 80,"Y": 30},"Ball":{"Pos":{"X": 5,"Y": 15	},"Vel": {"X": 0,"Y": 0}},"Hole": {"X": 75,"Y": 15}}'

function maketrigeometry(p1, p2, p3) {
    geo = new THREE.Geometry();
    geo.vertices.push(new THREE.Vector3(p1[0], p1[1], p1[2]));
    geo.vertices.push(new THREE.Vector3(p2[0], p2[1], p2[2]));
    geo.vertices.push(new THREE.Vector3(p3[0], p3[1], p3[2]));
    geo.faces.push(new THREE.Face3(0, 2, 1));
    geo.faces.push(new THREE.Face3(0, 1, 2));
    return geo;
}

function makewallgeometry(p1, p2, height) {
    p1x = p1[0];
    p1y = p1[1];
    p2x = p2[0];
    p2y = p2[1];

    geo1 = maketrigeometry([p1x, 0, p1y], [p2x, 0, p2y], [p1x, height, p1y]);
    geo2 = maketrigeometry([p2x, height, p2y], [p1x, height, p1y], [p2x, 0, p2y]);

    nmesh = new THREE.Mesh(geo1, red);
    nmesh.add(new THREE.Mesh(geo2, red));
    return nmesh;
}

function init() {

    redwireframe = new THREE.MeshBasicMaterial({
        color: 0xff0000,
        wireframe: true
    });

    red = new THREE.MeshBasicMaterial({
        color: 0xff0000
    });

    green = new THREE.MeshBasicMaterial({
        color: 0x00ff00
    });

    blue = new THREE.MeshBasicMaterial({
      color: 0xff //blue balls tehe tehehehe
    });

    black = new THREE.MeshBasicMaterial({
      color: 0x0
    })

    level = JSON.parse(levelstr);

    camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 1, 10000);
    camera.position.z = 100;
    camera.position.y = 100;
    camera.rotation.x = -3.141 / 4;

    scene = new THREE.Scene();

    ground = new THREE.Mesh(new THREE.PlaneGeometry(level.Dim.X, level.Dim.Y), green);
    ground.rotation.x = -3.141 / 2;
    ground.position.x = level.Dim.X/2;
    ground.position.z = level.Dim.Y/2;

    ball = new THREE.Mesh(new THREE.SphereGeometry(2.5), blue);
    ball.position.x = level.Ball.Pos.X;
    ball.position.z = level.Ball.Pos.Y;
    ball.position.y = 2.5;

    hole = new THREE.Mesh(new THREE.CircleGeometry(3), black);
    hole.position.x = level.Hole.X;
    hole.position.z = level.Hole.Y;
    hole.position.y += 0.01;
    hole.rotation.x = -3.141/2;

    scene.add(ground);
    scene.add(ball);
    scene.add(hole);
    scene.add(makewallgeometry([0,0],[level.Dim.X,0],5));
    scene.add(makewallgeometry([0,0],[0,level.Dim.Y],5));
    scene.add(makewallgeometry([level.Dim.X,0],[level.Dim.X,level.Dim.Y],5));
    scene.add(makewallgeometry([0,level.Dim.Y],[level.Dim.X,level.Dim.Y],5));

    renderer = new THREE.WebGLRenderer();
    renderer.setSize(window.innerWidth, window.innerHeight);

    document.body.appendChild(renderer.domElement);

}

function animate() {

    requestAnimationFrame(animate);
    scene.rotation.y += 0.01;
    renderer.render(scene, camera);

}
