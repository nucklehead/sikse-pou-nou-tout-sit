# Ripple

Effect that spreads like a wave in touch or click.
Inspired by Googles Material Design Guidelines.


## How To Use:

1. Download and include needed files:

```
	<link rel="stylesheet" href="dist/css/ripple.min.css" />
```

```
	<script src="dist/js/ripple.min.js"></script>
```

2. Add the `ripple` class to your HTML Element

```
	<button class="btn btn-success ripple">Ripple Button</button>
```

3. Enjoy.


## Customization

The default ripple style will create a white wave on click if you want to inverse that, just add `ripple-dark` as another class:

```
	<button class="btn btn-default ripple ripple-dark"> Ripple Dark Button</button>
```

you can add your own style by simply overwriting the `.ripple .ink` background property.

## SCSS Variables

you can simply override the following variables with your values:

```
$ripple-color:      #fff;
$ripple-dark-color: darken(#ccc, 30%);
$ripple-duration:   .5s;
```

[more Info and Examples](http://dev.raphaelmutschler.de/ripple/)
