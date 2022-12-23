

CREATE TABLE `foods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `url_image` varchar(255) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `date_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `date_updated` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Volcado de datos para la tabla `foods`
--

INSERT INTO `foods` (`id`, `name`, `category`, `description`, `url_image`, `status`, `date_created`, `date_updated`) VALUES
(1, 'pizza', 'fast', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://cdn.pixabay.com/photo/2017/12/09/08/18/pizza-3007395__480.jpg', 1, '2022-12-22 20:21:03', '2022-12-23 05:14:50'),
(2, 'hamburger', 'fast', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://cdn.pixabay.com/photo/2021/03/19/21/52/burger-6108495__480.jpg', 1, '2022-12-22 20:21:50', '2022-12-23 05:14:56'),
(3, 'tacos', 'mexican', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://cdn.pixabay.com/photo/2017/12/27/04/28/tortilla-3041938__480.jpg', 1, '2022-12-22 20:22:15', '2022-12-23 05:14:56'),
(4, 'pozole', 'mexican', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://media.istockphoto.com/id/915675212/es/foto/pozole-con-ma%C3%ADz-gran-mote-guisado-de-m%C3%A9xico.jpg?b=1&s=170667a&w=0&k=20&c=edXakG9VmwjFxgF1LY-luMCSS5kIecXq-84ww3NkT6Y=', 1, '2022-12-22 22:12:49', '2022-12-23 05:14:56'),
(5, 'memelas', 'mexican', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://media.istockphoto.com/id/1224503202/es/foto/sopes-mexicanos-con-queso-rallado-y-salsa-verde-comida-mexicana-picante-en-m%C3%A9xico.jpg?b=1&s=170667a&w=0&k=20&c=r1g04rIBPD3dohbnoIdV98cAscLiMKB4uKi70V1RaSE=', 1, '2022-12-22 20:22:15', '2022-12-23 05:14:56'),
(6, 'torta milanesa', 'mexican', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://media.istockphoto.com/id/1144414066/es/foto/comida-mexicana-s%C3%A1ndwich-de-chorizo-torta.jpg?b=1&s=170667a&w=0&k=20&c=KwJK_TzbiG5mh7M-GLbMUBFztJtNQT84tmYQQQfAf5A=', 1, '2022-12-22 20:22:15', '2022-12-23 05:14:56');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `foods`
--
ALTER TABLE `foods`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `foods`
--
ALTER TABLE `foods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
