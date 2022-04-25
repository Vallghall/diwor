import React from 'react'
import ProfileTable from "../ProfileTable/ProfileTable";

const DigestTable = ({ digests }) => {
    return (
        <table>
            <tbody>
            <tr><th colSpan="3">Ваши последние эксперименты</th></tr>
            <tr>
                <th>№</th>
                <th>Тип алгоритма</th>
                <th>Время начала</th>
            </tr>
            {<ProfileTable digests={digests}/>}
            </tbody>
        </table>
    )
}

export default DigestTable
