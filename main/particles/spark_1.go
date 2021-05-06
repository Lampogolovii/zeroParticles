components {
  id: "move"
  component: "/zeroParticles/particles/physics/zpPhysicsMove.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "gravity"
    value: "0.0, -500.0, 0.0"
    type: PROPERTY_TYPE_VECTOR3
  }
  properties {
    id: "initMin"
    value: "200.0, 150.0, 0.0"
    type: PROPERTY_TYPE_VECTOR3
  }
  properties {
    id: "initMax"
    value: "280.0, 200.0, 0.0"
    type: PROPERTY_TYPE_VECTOR3
  }
  properties {
    id: "velRotate"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
components {
  id: "delete"
  component: "/zeroParticles/particles/zpDeleteDuration.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
components {
  id: "alpha"
  component: "/zeroParticles/particles/color/zpSpriteAlphaFadeInOut.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "inDuration"
    value: "0.2"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "border"
  component: "/zeroParticles/particles/physics/zpPhysicsBorderHorizontal.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "minY"
    value: "-10.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "maxY"
    value: "-20.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "borderFriction"
    value: "0.3"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/images.atlas\"\n"
  "default_animation: \"spark_2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
