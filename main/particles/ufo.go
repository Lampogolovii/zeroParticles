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
  properties {
    id: "lifeDuration"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
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
    value: "0.1"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "pauseDuration"
    value: "0.7"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "outDuration"
    value: "0.2"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "outValue"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "trail"
  component: "/zeroParticles/particles/trail/zpTrailSin.script"
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
    id: "target"
    value: "310.0, 260.0, 0.0"
    type: PROPERTY_TYPE_VECTOR3
  }
  properties {
    id: "amplitudeUnit"
    value: "0.2"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "amplitudeUnitRandom"
    value: "0.1"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "sinWavesCount"
    value: "2.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "scale"
  component: "/zeroParticles/particles/transform/zpScaleTweenLinear.script"
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
    id: "scaleX"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "scaleY"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "startValue"
    value: "0.2"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "inDuration"
    value: "0.5"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "pauseDuration"
    value: "0.4"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "outDuration"
    value: "0.1"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "outValue"
    value: "0.5"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/images.atlas\"\n"
  "default_animation: \"ufo\"\n"
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
