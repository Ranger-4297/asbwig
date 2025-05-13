document.querySelectorAll('.saveRolesButton').forEach(button => {
	button.addEventListener('click', function () {
		const action = button.getAttribute('data-action');
		const inputId = `requiredRoles${action}Input`;
		const input = document.getElementById(inputId);
		
		const selectedRoles = Array.from(
			button.closest('.input-group').querySelectorAll('.roleCheckbox:checked')
		).map(cb => cb.value);
		
		input.value = JSON.stringify(selectedRoles);
		
		const baseurl = document.body.getAttribute('data-url');
		const guildID = document.body.getAttribute('data-guild-id');
		const url = `${baseurl}/dashboard/${guildID}/manage/update-roles`;
		
		fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ action: action, roles: selectedRoles }),
		}).then(() => {
			const alert = document.getElementById('successAlert');
			if (alert) {
				alert.style.display = 'block';
				alert.style.opacity = 1;
				setTimeout(() => {
					location.reload();
				}, 1500);
			}
		});
	}).catch(console.error);
});
