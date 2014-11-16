function messagein(data){
    parsed = JSON.parse(data);

    //If this is the message with the geometry, start
    if(parsed["leveldata"] != null){
      init(parsed.leveldata);
    }
}

function maketrigeometry(p1, p2, p3) {
  geo = new THREE.Geometry();
  geo.vertices.push(new THREE.Vector3(p1[0], p1[1], p1[2]));
  geo.vertices.push(new THREE.Vector3(p2[0], p2[1], p2[2]));
  geo.vertices.push(new THREE.Vector3(p3[0], p3[1], p3[2]));
  geo.faces.push(new THREE.Face3(0, 2, 1));
  geo.faces.push(new THREE.Face3(0, 1, 2));
  return geo;
}

function MakeWall(p1, p2, height) {
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

function init(data) {

  redwireframe = new THREE.MeshBasicMaterial({
    color: 0xff0000,
    wireframe: true
  });

  white = new THREE.MeshBasicMaterial({
    color: 0xffffff
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

  level = data;

  scene = new THREE.Scene();

  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 1, 10000);
  camera.position.z = 100;
  camera.position.y = 100;
  camera.rotation.x = -3.141 / 4;

  gplane = new THREE.Plane(new THREE.Vector3(0,-1,0), 0);

  ground = new THREE.Mesh(new THREE.PlaneGeometry(level.Dim.X, level.Dim.Y), green);
  ground.rotation.x = -3.141 / 2;
  ground.position.x = level.Dim.X / 2;
  ground.position.z = level.Dim.Y / 2;

  ball = new THREE.Mesh(new THREE.SphereGeometry(2.5), blue);
  ball.position.x = level.Ball.Pos.X;
  ball.position.z = level.Ball.Pos.Y;
  ball.position.y = 2.5;

  hole = new THREE.Mesh(new THREE.CircleGeometry(3), black);
  hole.position.x = level.Hole.X;
  hole.position.z = level.Hole.Y;
  hole.position.y += 0.01;
  hole.rotation.x = -3.141 / 2;


  linegeo = new THREE.Geometry();
  linegeo.vertices.push(ball.position);
  linegeo.vertices.push(hole.position);
  shootline = new THREE.Line(linegeo, white);

  scene.add(ground);
  scene.add(ball);
  scene.add(hole);
  scene.add(shootline);
  scene.add(MakeWall([0, 0], [level.Dim.X, 0], 5));
  scene.add(MakeWall([0, 0], [0, level.Dim.Y], 5));
  scene.add(MakeWall([level.Dim.X, 0], [level.Dim.X, level.Dim.Y], 5));
  scene.add(MakeWall([0, level.Dim.Y], [level.Dim.X, level.Dim.Y], 5));

  renderer = new THREE.WebGLRenderer();
  renderer.setSize(window.innerWidth, window.innerHeight);

  document.body.appendChild(renderer.domElement);
  animate();
}

function animate() {
  requestAnimationFrame(animate);
  //scene.rotation.y += 0.0025;
  renderer.render(scene, camera);
}
