document.getElementById('saveButton').addEventListener('click', function() {
    const input = document.getElementById('prefixInput').value.trim();
    if (input === '') {
        document.getElementById('dangerAlert').style.display = 'block';
        document.getElementById('dangerAlert').style.opacity = 1;
        document.getElementById('successAlert').style.display = 'none';
    } else {
        document.getElementById('successAlert').style.display = 'block';
        document.getElementById('successAlert').style.opacity = 1;
        document.getElementById('dangerAlert').style.display = 'none';
        const baseurl = document.body.getAttribute('data-url');
        const guildID = document.body.getAttribute('data-guild-id');
        const url = `${baseurl}/dashboard/${guildID}/manage/update-prefix`;
        console.log(input);
        fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ prefix: input }),
        }).then(() => {
			setTimeout(() => {
				location.reload();
			}, 1500);
        }).catch(console.error);
    }

    setTimeout(function() {
        document.getElementById('dangerAlert').style.opacity = 0;
        setTimeout(function() {
            document.getElementById('dangerAlert').style.display = 'none';
        }, 300);

        document.getElementById('successAlert').style.opacity = 0;
        setTimeout(function() {
            document.getElementById('successAlert').style.display = 'none';
        }, 300);
    }, 3000);
});