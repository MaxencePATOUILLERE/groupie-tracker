function calculateLongLat(location, callback) {
    // utilisation de l'API de géocodage de Google Maps pour obtenir les coordonnées géographiques de la ville
    var geocoder = new google.maps.Geocoder();
    var latLng = { lat: 0, lng: 0 };
    geocoder.geocode({ address: location }, function (results, status) {
        if (status == google.maps.GeocoderStatus.OK) {
            latLng = results[0].geometry.location;
            callback(latLng);
        }
    });
}